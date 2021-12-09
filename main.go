package main

import (
	"api/golang/internal/http"
	"log"
)

func main() {
	log.Println("Starting server 🚀")

	error := http.StartServer()
	log.Fatal(error)
}
