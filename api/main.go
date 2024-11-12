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
	fmt.Println(fmt.Sprintf("Executando a api na porta: %d", config.Porta))
	fmt.Println(fmt.Sprintf("String de conexão com o banco: %s", config.StringConexaoBanco))

	router := router.Gerar()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), router))
}
