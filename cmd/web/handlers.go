package main

import "net/http"

// VirtualTerminal é um manipulador HTTP que será chamado quando a rota
// "/virtual-terminal" for acessada
func (app *application) VirtualTerminal(writer http.ResponseWriter, request *http.Request) {
	if err := app.renderTemplate(writer, request, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}
