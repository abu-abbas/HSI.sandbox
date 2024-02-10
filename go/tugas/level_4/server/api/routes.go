package api

import (
	"net/http"

	"github.com/abu-abbas/level_4/server/controllers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func (app *AppServer) routes() {
	app.router.Use(middleware.RequestID)
	app.router.Use(middleware.Logger)
	app.router.Use(middleware.Recoverer)
	app.router.Use(middleware.URLFormat)
	app.router.Use(render.SetContentType(render.ContentTypeJSON))
	app.router.Route(
		"/api",
		func(r chi.Router) {
			r.Mount("/v1", v1Handler(r))
		},
	)

	// base path
	app.router.Get(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("root."))
		},
	)
}

func v1Handler(r chi.Router) http.Handler {
	r.Mount("/items", itemServices(r))

	return r
}

func itemServices(r chi.Router) http.Handler {
	controllerItem := controllers.Item{}

	r.Get("/", controllerItem.Index)
	r.Post("/", controllerItem.Create)
	r.Route("/{id}", func(r chi.Router) {
		r.Put("/", controllerItem.Edit)
		r.Delete("/", controllerItem.Delete)
	})

	return r
}
