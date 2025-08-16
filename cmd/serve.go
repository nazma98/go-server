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
		middleware.Hudai,
		middleware.Logger,
	)(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /route", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	mux.Handle("POST /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.CreateProduct))))
	mux.Handle("GET /products/{id}", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductByID))))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
