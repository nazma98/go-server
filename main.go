package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Nazma")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {

	sendData(w, productList, 200)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "PLease give me valid json", 400)
		return
	}

	newProduct.ID = len(productList) + 1

	productList = append(productList, newProduct)

	sendData(w, newProduct, 201)
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /hello", http.HandlerFunc(helloHandler))

	mux.Handle("GET /about", http.HandlerFunc(aboutHandler))

	mux.Handle("GET /products", http.HandlerFunc(getProducts))

	mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))

	mux.Handle("POST /create-products", http.HandlerFunc(createProduct))

	mux.Handle("OPTIONS /create-products", http.HandlerFunc(createProduct))

	fmt.Println("Server running on :8080")

	globalRouter := globalRouter(mux)

	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	prd1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is orange. I love orange.",
		Price:       100,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTct4QnA6zwU4tyn9vq_ADJ8Ey6RfIWK-6V6g&s",
	}

	prd2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green. I don't really like apple.",
		Price:       180,
		ImgURL:      "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg?v=1623908361&width=1920",
	}

	prd3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow. I feel bored eating banana.",
		Price:       5,
		ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}

	// prd4 := Product{
	// 	ID:          3,
	// 	Title:       "Orange",
	// 	Description: "Orange is orange. I love orange.",
	// 	Price:       100,
	// 	ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTct4QnA6zwU4tyn9vq_ADJ8Ey6RfIWK-6V6g&s",
	// }

	// prd5 := Product{
	// 	ID:          4,
	// 	Title:       "Apple",
	// 	Description: "Apple is green. I don't really like apple.",
	// 	Price:       180,
	// 	ImgURL:      "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg?v=1623908361&width=1920",
	// }

	// prd6 := Product{
	// 	ID:          5,
	// 	Title:       "Mango",
	// 	Description: "Mango is yummy. I love it very much.",
	// 	Price:       4565,
	// 	ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	// }

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	// productList = append(productList, prd4)
	// productList = append(productList, prd5)
	// productList = append(productList, prd6)

}

func globalRouter(mux *http.ServeMux) http.Handler {
	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Nazma")
		w.Header().Set("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}
		mux.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handleAllReq)
}
