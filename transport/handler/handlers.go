package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := []string{"Phone", "Laptop", "Headphones", "Sex toys"}
	respondWithJSON(w, http.StatusOK, products)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	// Получение динамического параметра из URL через chi
	productID := chi.URLParam(r, "id")
	respondWithJSON(w, http.StatusOK, map[string]string{
		"id":   productID,
		"name": "Laptop",
	})
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob"}
	respondWithJSON(w, http.StatusOK, users)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	respondWithJSON(w, http.StatusOK, map[string]string{
		"id":   userID,
		"name": "Alice",
	})
}

func GistUserOrdersHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	orders := []string{"Order_1", "Order_2"}
	respondWithJSON(w, http.StatusOK, map[string]string{
		"user_id": userID,
		"orders":  fmt.Sprintf("%v", orders),
	})
}

func ListUserOrdersHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	orders := []string{"Order_1", "Order_2"}
	respondWithJSON(w, http.StatusOK, map[string]string{
		"user_id": userID,
		"orders":  fmt.Sprintf("%v", orders),
	})
}

// Хелпер для быстрой отправки JSON-ответов
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
