package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	stringConexaoBanco = ""
	porta              = 0
)

// Carregar - Carrega as configurações do sistema - variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	porta, erro = strconv.Atoi(os.Getenv("API_PORTA"))
	if erro != nil {
		porta = 9000
	}

	stringConexaoBanco = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_NOME"),
	)
}

// StringConexaoBanco - Retorna a string de conexão com o banco de dados
func StringConexaoBanco() string {
	return stringConexaoBanco
}

// Porta - Retorna a porta que a API irá subir
func Porta() int {
	return porta
}
