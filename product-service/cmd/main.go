package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mangesh-shinde/ecomm-app/product-service/internal/config"
	"github.com/mangesh-shinde/ecomm-app/product-service/internal/http/handlers"
)

func main() {

	// load config file
	cfg := config.MustLoad()

	// setup database
	// db := storage.Connect(cfg)

	// setup http server
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.Mount("/api/products", handlers.ProductRoutes())

	fmt.Printf("Server is listerning on address: %s\n", cfg.Addr)
	log.Fatal(server.ListenAndServe())

}
