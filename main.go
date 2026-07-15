package main

import (
	"fmt"
	"net/http"
	"rest-demo/transport/handler"    // --- ОБРАБОТЧИКИ (HANDLERS) ---
	"rest-demo/transport/middleware" // --- МИДЛВАРИ ---
	"time"

	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// 1. Глобальные мидлвари (применяются ко ВСЕМ запросам на сервере)
	r.Use(chimiddleware.Logger)    // Логирование запросов в консоль
	r.Use(chimiddleware.Recoverer) // Защита от паники (panic recovery)
	r.Use(chimiddleware.Timeout(60 * time.Second))

	// 2. Глобальная группировка API (Версионирование)
	r.Route("/api/v1", func(r chi.Router) {

		// 3. Саброутер для ТОВАРОВ (Публичный)
		// Все роуты внутри начинаются с префикса "/api/v1/products"
		r.Route("/products", func(r chi.Router) {
			r.Get("/", handler.ListProductsHandler)   // GET /api/v1/products
			r.Get("/{id}", handler.GetProductHandler) // GET /api/v1/products/123
		})

		// 4. Саброутер для ПОЛЬЗОВАТЕЛЕЙ (Приватный)
		// Все роуты внутри начинаются с префикса "/api/v1/users"
		r.Route("/users", func(r chi.Router) {
			// Локальная мидлварь — применяется только к роутам внутри группы "/users"
			r.Use(middleware.AuthMiddleware)

			r.Get("/", handler.ListUsersHandler)   // GET /api/v1/users
			r.Get("/{id}", handler.GetUserHandler) // GET /api/v1/users/45

			// 5. Глубоко вложенный саброутер для ЗАКАЗОВ пользователя
			// Полный путь: "/api/v1/users/{id}/orders"
			r.Route("/{id}/orders", func(r chi.Router) {
				r.Get("/", handler.ListUserOrdersHandler) // GET /api/v1/users/45/orders
			})
		})
	})

	fmt.Println("REST API Server running on :8080")
	http.ListenAndServe(":8080", r)
}
