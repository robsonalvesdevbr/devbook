package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MySql Driver
)

// Conectar - Retorna uma conexão com o banco de dados
func Conectar() (*sql.DB, error) {
	conexao, erro := sql.Open("mysql", config.StringConexaoBanco())
	if erro != nil {
		return nil, erro
	}

	if erro = conexao.Ping(); erro != nil {
		conexao.Close()
		return nil, erro
	}

	return conexao, nil
}
