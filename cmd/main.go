package main

import (
	"encurtador-url/internal/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/shorten", handler.ShortenHandler)
	http.HandleFunc("/", handler.RedirectHandler)

	log.Println("Servidor escutando na porta 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
