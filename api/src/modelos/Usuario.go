package modelos

import (
	"errors"
	"strings"
	"time"
)

// Usuario - Estrutura que representa um usuário do sistema
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criado_em,omitempty"`
}

// validar - Método para chamar os métodos de validação e formatação do usuário
func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}

	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}

	if usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}

	return nil
}

// formatar - Método para chamar o método de formatação do usuário
func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}

// Preparar - Método para chamar os métodos de formatação e validação do usuário
func (usuario *Usuario) Preparar() error {
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	usuario.formatar()
	return nil
}
