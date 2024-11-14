package servidor

import (
	"fmt"
	"net/http"
)

// Servidor - Interface para iniciar o servidor
type Servidor interface {
	Iniciar() error
}

// servidorHTTP - Implementação do servidor HTTP
type servidorHTTP struct {
	porta  int
	router http.Handler
}

// NovoServidorHTTP - Cria uma nova instância do servidor HTTP
func NovoServidorHTTP(porta int, router http.Handler) Servidor {
	return &servidorHTTP{porta: porta, router: router}
}

// Iniciar - Inicia o servidor HTTP
func (s *servidorHTTP) Iniciar() error {
	endereco := fmt.Sprintf(":%d", s.porta)
	fmt.Printf("Executando a API na porta: %d\n", s.porta)
	return http.ListenAndServe(endereco, s.router)
}
