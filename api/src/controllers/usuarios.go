package controllers

import "net/http"

// CriarUsuario cria um usuário
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando usuário"))
}

// BuscarUsuarios busca todos os usuários
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuários"))
}

// BuscarUsuario busca um usuário
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar usuário"))
}

// AtualizarUsuario atualiza um usuário
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizar usuário"))
}

// DeletarUsuario deleta um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletar usuário"))
}
