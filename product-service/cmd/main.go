package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mangesh-shinde/ecomm-app/product-service/internal/config"
)

func main() {

	// load config file
	cfg := config.MustLoad()

	// setup http server
	router := chi.NewRouter()

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	fmt.Printf("Server is listerning on address: %s\n", cfg.Addr)
	log.Fatal(server.ListenAndServe())

}
