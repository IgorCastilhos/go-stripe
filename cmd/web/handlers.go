package main

import "net/http"

// VirtualTerminal é um manipulador HTTP que será chamado quando a rota
// "/virtual-terminal" for acessada
func (app *application) VirtualTerminal(writer http.ResponseWriter, request *http.Request) {
	// Registra uma mensagem no log de informações indicando que o manipulador foi atingido
	app.infoLog.Println("Atingiu o manipulador")
}
