package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// routes Método cria e retorna um manipulador HTTP (http.Handler)
// configurado com rotas específicas
func (app *application) routes() http.Handler {
	// Cria um novo roteador multiplexer usando o pacote chi
	mux := chi.NewRouter()

	// Adiciona uma rota para a URL "/virtual-terminal" que será manipulada
	// pelo método VirtualTerminal da aplicação
	mux.Get("/virtual-terminal", app.VirtualTerminal)

	// Retorna o roteador configurado como um manipulador HTTP
	return mux
}
