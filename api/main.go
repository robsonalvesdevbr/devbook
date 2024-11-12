package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	fmt.Printf("Executando a api na porta: %d\n", config.Porta())
	fmt.Printf("String de conexão com o banco: %s\n", config.StringConexaoBanco())

	router := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta()), router))
}
