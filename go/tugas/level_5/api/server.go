package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/abu-abbas/level_5/config"
	"github.com/gofiber/fiber/v2"
)

type AppServer struct {
	config config.ServerConfig
	app    *fiber.App
}

func CreateApp() *AppServer {
	svr := &AppServer{
		config: config.GetYamlValue().ServerConfig,
		app:    fiber.New(),
	}

	svr.routes()
	return svr
}

func (server *AppServer) Mount() {
	ctx := context.Background()
	shutdownComplete := handleShutdown(func() {
		if err := server.app.ShutdownWithContext(ctx); err != nil {
			log.Printf("server shutdown failed: %v\n", err)
		}
	})

	log.Printf("http listen and serve on port: %s", server.config.Port)
	if err := server.app.Listen(fmt.Sprintf(":%s", server.config.Port)); err == http.ErrServerClosed {
		<-shutdownComplete
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
