/*
	Este código organiza a renderização de templates HTML
*/

package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// templateData definição da estrutura para armazenar dados usados nos templates
type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float64
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

// Declaração do map de funções para templates (vazio neste caso)
var functions = template.FuncMap{}

// Incorpora os templates no binário durante a compilação
//
//go:embed templates
var templateFS embed.FS

// addDefaultData Método para adicionar dados padrões à estrutura templateData
func (app *application) addDefaultData(td *templateData, request *http.Request) *templateData {
	return td
}

// renderTemplate Método para renderizar templates HTML
func (app *application) renderTemplate(writer http.ResponseWriter, request *http.Request, page string, td *templateData, partials ...string) error {
	var t *template.Template
	var err error
	templateToRender := fmt.Sprintf("templates/%s.page.tmpl", page)

	// Verifica se o template está no cache antes de tentar reconstruir
	_, templateInMap := app.templateCache[templateToRender]

	// Se estiver em modo "production" e o template estiver no cache, usa o cache, senão reconstrói
	if app.config.env == "production" && templateInMap {
		t = app.templateCache[templateToRender]
	} else {
		t, err = app.parseTemplate(partials, page, templateToRender)
		if err != nil {
			app.errorLog.Println(err)
			return err
		}
	}

	// Se a estrutura templateData for nula, cria uma nova
	if td == nil {
		td = &templateData{}
	}

	// Adiciona dados padrão à estrutura templateData
	td = app.addDefaultData(td, request)

	// Executa o template com os dados e escreve no ResponseWrite
	err = t.Execute(writer, td)
	if err != nil {
		app.errorLog.Println(err)
		return err
	}

	return nil
}

// parseTemplate Método para analisar e criar um template
func (app *application) parseTemplate(partials []string, page, templateToRender string) (*template.Template, error) {
	var t *template.Template
	var err error

	// Constrói os caminhos para os partials
	if len(partials) > 0 {
		for i, x := range partials {
			partials[i] = fmt.Sprintf("templates/%s.partial.tmpl", x)
		}
	}

	// Cria o template com os partials e o template principal
	if len(partials) > 0 {
		t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.tmpl", strings.Join(partials, ","), templateToRender)
	} else {
		// Caso não haja partials
		t, err = template.New(fmt.Sprintf("%s.page.tmpl", page)).Funcs(functions).ParseFS(templateFS, "templates/base.layout.tmpl", templateToRender)
	}
	if err != nil {
		app.errorLog.Println(err)
		return nil, err
	}

	// Adiciona o template ao cache
	app.templateCache[templateToRender] = t
	return t, nil
}
