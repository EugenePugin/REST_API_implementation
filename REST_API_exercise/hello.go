package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// 1. Глобальные мидлвари (применяются ко ВСЕМ запросам на сервере)
	r.Use(middleware.Logger)    // Логирование запросов в консоль
	r.Use(middleware.Recoverer) // Защита от паники (panic recovery)
	r.Use(middleware.Timeout(60 * time.Second))

	// 2. Глобальная группировка API (Версионирование)
	r.Route("/api/v1", func(r chi.Router) {

		// 3. Саброутер для ТОВАРОВ (Публичный)
		// Все роуты внутри начинаются с префикса "/api/v1/products"
		r.Route("/products", func(r chi.Router) {
			r.Get("/", listProductsHandler)   // GET /api/v1/products
			r.Get("/{id}", getProductHandler) // GET /api/v1/products/123
		})

		// 4. Саброутер для ПОЛЬЗОВАТЕЛЕЙ (Приватный)
		// Все роуты внутри начинаются с префикса "/api/v1/users"
		r.Route("/users", func(r chi.Router) {
			// Локальная мидлварь — применяется только к роутам внутри группы "/users"
			r.Use(authMiddleware)

			r.Get("/", listUsersHandler)   // GET /api/v1/users
			r.Get("/{id}", getUserHandler) // GET /api/v1/users/45

			// 5. Глубоко вложенный саброутер для ЗАКАЗОВ пользователя
			// Полный путь: "/api/v1/users/{id}/orders"
			r.Route("/{id}/orders", func(r chi.Router) {
				r.Get("/", listUserOrdersHandler) // GET /api/v1/users/45/orders
			})
		})
	})

	fmt.Println("REST API Server running on :8080")
	http.ListenAndServe(":8080", r)
}

// --- МИДЛВАРИ ---

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "secret-token" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

// --- ОБРАБОТЧИКИ (HANDLERS) ---

func listProductsHandler(w http.ResponseWriter, r *http.Request) {
	products := []string{"Phone", "Laptop", "Headphones", "Sex toys"}
	respondWithJSON(w, http.StatusOK, products)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	// Получение динамического параметра из URL через chi
	productID := chi.URLParam(r, "id")
	respondWithJSON(w, http.StatusOK, map[string]string{
		"id":   productID,
		"name": "Laptop",
	})
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []string{"Alice", "Bob"}
	respondWithJSON(w, http.StatusOK, users)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")
	respondWithJSON(w, http.StatusOK, map[string]string{
		"id":   userID,
		"name": "Alice",
	})
}

func listUserOrdersHandler(w http.ResponseWriter, r *http.Request) {
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
