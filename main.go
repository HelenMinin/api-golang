package main

import (
	"facec/blog/internal/http"
	"log"
)

func main() {
	log.Println("iniciando...")

	err := http.StartServer()
	log.Fatalln(err)
}
