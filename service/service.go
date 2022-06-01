package service

import (
	"activeReconBot/service/config"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize(config *config.Config) {
	//dbURI := fmt.Sprintf("%s://%s:%s@%s:%d",
	//	config.DB.Dialect,
	//	config.DB.Username,
	//	config.DB.Password,
	//	config.DB.Host,
	//	config.DB.Port)
	dbURI := fmt.Sprintf("%s://%s:%d",
		config.DB.Dialect,
		config.DB.Host,
		config.DB.Port)

	err := mgm.SetDefaultConfig(nil, config.DB.Name, options.Client().ApplyURI(dbURI))
	if err != nil {
		log.Fatalln(err)
	}

	//
	//a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func corsHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			//handle preflight in here
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")

		} else {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.Header().Add("Access-Control-Allow-Headers", "*")
			h.ServeHTTP(w, r)
		}
	}
}

// setRouters sets the all required routers
func (a *App) setRouters() {

	//search test
	//a.Get("/test", corsHandler(a.handleRequest(handler.Test)))

}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST", "OPTIONS")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("DELETE")
}

// Run the app on it's router
func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *App) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
