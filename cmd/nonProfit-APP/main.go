package main

import (
	"log"
	"net/http"

	"github.com/teremuhamblin/nonProfit-APP/internal/http"
)

func main() {
	router := http.NewRouter()

	log.Println("NonProfitApp Go server running on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
