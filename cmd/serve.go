package cmd

import (
	"fmt"
	"net/http"

	"ecommerce/middleware"
)

func Serve() {
	manager := middleware.NewManager()

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	wrappedMux := manager.WrapMux(
		mux,
		middleware.Logger,
		middleware.Hudai,
		middleware.CorsWithPreflight)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}
