package database

var ProductList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
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

	prd4 := Product{
		ID:          3,
		Title:       "Orange",
		Description: "Orange is orange. I love orange.",
		Price:       100,
		ImgURL:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTct4QnA6zwU4tyn9vq_ADJ8Ey6RfIWK-6V6g&s",
	}

	prd5 := Product{
		ID:          4,
		Title:       "Apple",
		Description: "Apple is green. I don't really like apple.",
		Price:       180,
		ImgURL:      "https://www.harrisfarm.com.au/cdn/shop/products/40715-done.jpg?v=1623908361&width=1920",
	}

	prd6 := Product{
		ID:          5,
		Title:       "Mango",
		Description: "Mango is yummy. I love it very much.",
		Price:       4565,
		ImgURL:      "https://www.allrecipes.com/thmb/lc7nSL9L5zMHXz9t6PMAVm9biNM=/1500x0/filters:no_upscale():max_bytes(150000):strip_icc()/ar-new-banana-adobe-ar-2x1-917fdde58d194b529b41042ebff1c031.jpg",
	}

	ProductList = append(ProductList, prd1)
	ProductList = append(ProductList, prd2)
	ProductList = append(ProductList, prd3)
	ProductList = append(ProductList, prd4)
	ProductList = append(ProductList, prd5)
	ProductList = append(ProductList, prd6)

}
