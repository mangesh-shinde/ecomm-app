package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mangesh-shinde/ecomm-app/product-service/internal/storage"
	"github.com/mangesh-shinde/ecomm-app/product-service/pkg/models"
	"github.com/mangesh-shinde/ecomm-app/product-service/pkg/utils"
)

func ProductRoutes() chi.Router {
	r := chi.NewRouter()

	p := ProductHandler{}
	r.Post("/", p.AddProduct())
	r.Get("/", p.GetProducts())
	r.Get("/{id}", p.GetProduct())
	r.Put("/{id}", p.UpdateProduct())
	r.Delete("/{id}", p.DeleteProduct())

	return r
}

type ProductHandler struct {
	Store storage.Storage
}

func (p *ProductHandler) GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		response := map[string]string{"message": "All products list here"}
		json.NewEncoder(w).Encode(response)
	}
}

func (p *ProductHandler) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		productId := chi.URLParam(r, "id")
		response := map[string]string{"message": fmt.Sprintf("Product with id=%s details here", productId)}
		json.NewEncoder(w).Encode(response)
	}
}

func (p *ProductHandler) AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// get product details through Request
		var product models.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusBadRequest, err)
			return
		}
		fmt.Printf("product: %+v\n", product)
		// add document to mongodb
		// p.Store.Add()

		// send response to client
		utils.WriteJsonResponse(w, http.StatusCreated, map[string]string{"message": "Product added"})
	}
}

func (p *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		productId := chi.URLParam(r, "id")
		response := map[string]string{"message": fmt.Sprintf("Product %s details updated", productId)}
		json.NewEncoder(w).Encode(response)
	}
}

func (p *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		productId := chi.URLParam(r, "id")
		response := map[string]string{"message": fmt.Sprintf("Product %s removed", productId)}
		json.NewEncoder(w).Encode(response)
	}
}
