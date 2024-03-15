package main

import (
	"log"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/omarghetti/test_webserver/v2/api/v1"
	"github.com/omarghetti/test_webserver/v2/db"
)

type App struct {
	Router   *http.ServeMux
	Database *pg.DB
	Topics   []string
}

func (app *App) Initialize() {
	// Initialize the database
	// Initialize the router
	// Initialize the topics
	app.Database = db.ConnectToDatabase()
	app.Router.HandleFunc("GET /hello", api.PrintHelloWorld)
	app.Topics = []string{"topic1", "topic2", "topic3"}
}

func main() {
	app := &App{
		Router: http.NewServeMux(),
	}
	app.Initialize()
	log.Printf("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
