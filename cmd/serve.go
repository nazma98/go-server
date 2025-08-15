package cmd

import (
	"fmt"
	"log"
	"net/http"

	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
)

func Serve() {
	mux := http.NewServeMux()

	cntrl := func(w http.ResponseWriter, r *http.Request) {
		log.Println("ami hnadler, middle e print hobo")
	}

	handler := http.HandlerFunc(cntrl)

	mux.Handle("GET /route", middleware.Logger(handler))

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
