package main

import (
	"embed"
	"io/fs"
	"log"
	"log/slog"
	"net/http"
	"ordersystem/repository"
	"ordersystem/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	_ "ordersystem/docs"
	// OpenApi
	httpSwagger "github.com/swaggo/http-swagger"
)

// embeddedFrontend embeds the frontend into the binary
//
//go:embed frontend/*
var embeddedFrontend embed.FS

// @title				Order System
// @description			This system enables drink orders and should not be used for the forbidden Hungover Games.
// @contact.name		Your Name
func main() {
	repo := repository.Connect()
	repo.InitSchema()

	log.Println("🚀 Ordersystem API starting...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/swagger/*", httpSwagger.WrapHandler)
	// Render Frontend
	staticFS, err := fs.Sub(embeddedFrontend, "frontend")
	if err != nil {
		log.Fatal(err)
	}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFileFS(w, r, staticFS, "index.html")
	})
	// Menu Routes
	r.Get("/api/menu", rest.GetMenu(repo))
	// Order Routes
	r.Get("/api/order/all", rest.GetOrders(repo))
	r.Get("/api/order/totalled", rest.GetOrdersTotal(repo))
	r.Post("/api/order", rest.PostOrder(repo))
	// OpenAPI Routes
	r.Get("/openapi/*", httpSwagger.WrapHandler)

	slog.Info("⚡⚡⚡ Order System is up and buzzin ⚡⚡⚡")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
