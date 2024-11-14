package main

import (
	"api/src/config"
	"api/src/router"
	"api/src/servidor"
	"log"
)

func main() {
	config.Carregar()

	porta := config.Porta()
	rotas := router.Gerar()

	srv := servidor.NovoServidorHTTP(porta, rotas)
	if err := srv.Iniciar(); err != nil {
		log.Fatal(err)
	}
}
