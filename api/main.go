package main

import (
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Executando a api...")

	router := router.Gerar()
	log.Fatal(http.ListenAndServe(":5000", router))
}
