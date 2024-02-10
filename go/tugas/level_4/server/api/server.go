package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/abu-abbas/level_4/server/config"
	"github.com/go-chi/chi"
)

type AppServer struct {
	config config.ServerConfig
	router *chi.Mux
}

func CreateApp() *AppServer {
	cfg := config.GetYamlValue().ServerConfig
	svr := &AppServer{
		config: cfg,
		router: chi.NewRouter(),
	}

	svr.routes()

	return svr
}

func (app *AppServer) Mount() {
	ctx := context.Background()

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", app.config.Port),
		Handler: app.router,
	}

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server shutdown failed: %v\n", err)
		}
	})

	log.Printf("http listen and serve on port: %s", app.config.Port)
	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http listen and serve failed: %v\n", err)
	}

	log.Println("shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})
	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
