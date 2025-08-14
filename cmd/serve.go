package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/global_router"
	"ecommerce/handlers"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProduct))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
