/*
	Este código inclui configurações para o servidor, tratamento de logs, manipulação de templates e inicialização da aplicação.
*/

package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// cssVersion é incrementado para forçar a atualização de arquivos CSS e JS no cache do navegador,
// garantindo que os usuários recebam as versões mais recentes após alterações no código.
const cssVersion = "1"

// config armazena configurações para a aplicação
type config struct {
	port int
	// env = production, development por ex.
	env string
	api string
	db  struct {
		dsn string
	}
	stripe struct {
		secret string
		key    string
	}
}

// application - Recebedor da maioria das partes da aplicação
type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

func (app *application) serve() error {
	// Configuração do servidor HTTP
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	// Iniciando o servidor
	app.infoLog.Printf(fmt.Sprintf("Iniciando servidor HTTP em modo %s na porta %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	// Configuração das flags de linha de comando
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Porta para o servidor escutar")
	flag.StringVar(&cfg.env, "env", "development", "Ambiente da aplicação {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL para a api")
	flag.Parse()

	// Configuração das credenciais Stripe a partir de variáveis de ambiente
	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	// Configuração dos logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// tc é o cache de templates
	tc := make(map[string]*template.Template)

	// Inicialização da aplicação
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}

	// Inicia o servidor
	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
