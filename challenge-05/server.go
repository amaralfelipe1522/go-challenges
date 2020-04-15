package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/cart/", Orquestrador)
	log.Println("Executando na porta 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
