package main

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	"net/http"

	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/maciekmm/uek-bruschetta/controllers"
)

func main() {
	logger := log.New(os.Stdout, "Bruschette", log.Lshortfile)
	app := &Application{Logger: logger}

	err := app.init()
	if err != nil {
		logger.Fatal(err)
	}

	app.setupRoutes()

	err = app.serve()
	if err != nil {
		logger.Fatal(err)
	}
}

type Application struct {
	Database *sql.DB
	Logger   *log.Logger
	router   *mux.Router
}

func (a *Application) init() error {
	a.Logger.Println("starting Bruschette")

	a.Logger.Println("setting up database connection")
	con, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@database/%s?sslmode=disable", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		return fmt.Errorf("could not open database connection: %s", err.Error())
	}

	a.Logger.Println("establishing database connection")
	deadline := time.After(5 * time.Second)
out:
	for {
		select {
		case <-deadline:
			return fmt.Errorf("could not establish database connection, last error: %s", err.Error())
		default:
			err = con.Ping()
			if err == nil {
				break out
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

	return nil
}

func (a *Application) setupRoutes() {
	// setup routes
	a.Logger.Println("setting up routes")
	a.router = mux.NewRouter()
	accountController := &controllers.Account{}
	accountController.Register(a.router.PathPrefix("/account/").Subrouter())
}

func (a *Application) serve() error {
	return http.ListenAndServe(":3000", a.router)
}
