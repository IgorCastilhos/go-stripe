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

// cssVersion será adicionado a qualquer arquivo CSS ou JS. Quando incrementado, isso forçará a maioria dos navegadores
// a desativar a nova versão. Isso resolverá o problema de ter que limpar o cache quando as coisas não estiverem
// funcionando
const cssVersion = "1"

// config armazena configurações pra aplicação
type config struct {
	port int
	// env = production, development por ex.
	env string
	api string
	// db - dsn é o Data Source Name
	db struct {
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
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Iniciando servidor HTTP em modo %s na porta %d", app.config.env, app.config.port)
	return srv.ListenAndServe()
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Porta para o servidor escutar")
	flag.StringVar(&cfg.env, "env", "development", "Ambiente da aplicação {development|production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:40001", "URL para a api")

	flag.Parse()

	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// tc é o template cache
	tc := make(map[string]*template.Template)

	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}
	err := app.serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
