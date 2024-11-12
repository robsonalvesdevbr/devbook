package controllers

import (
	"api/src/modelos"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// CriarUsuario - Cria um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := io.ReadAll(r.Body)
	if erro == nil {
		log.Fatal(erro)
	}

	var usuario modelos.Usuario
	if erro := json.Unmarshal(corpoRequest, &usuario); erro != nil {
		log.Fatal(erro)
	}
}

// BuscarUsuarios - Busca todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuários"))
}

// BuscarUsuario - Busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuário"))
}

// AtualizarUsuario - Atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuário"))
}

// DeletarUsuario - Deleta um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuário"))
}
