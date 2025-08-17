package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	mux.Handle("GET /shimul", manager.With(
		http.HandlerFunc(handlers.Test),
		middleware.Logger,
		middleware.Hudai,
	))

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
		middleware.Logger,
		middleware.Hudai))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProduct),
		middleware.Logger,
		middleware.Hudai))

	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProductByID),
		middleware.Logger,
		middleware.Hudai))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
