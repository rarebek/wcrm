package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	http.HandleFunc("/products", handleProducts)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	// Define some sample products
	products := []Product{
		{Name: "Product 1", Description: "Description of Product 1", Price: 1000000},
		{Name: "Product 2", Description: "Description of Product 2", Price: 2000000},
		{Name: "Product 3", Description: "Description of Product 3", Price: 3000000},
	}

	// Marshal products to JSON
	jsonData, err := json.Marshal(products)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write JSON response
	_, err = w.Write(jsonData)
	if err != nil {
		log.Println("Error writing JSON response:", err)
	}
}
