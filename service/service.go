package service

import (
	"activeReconBot/service/config"
	"activeReconBot/service/handler"
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
	a.Router.Use(loggingMiddleWare)
	//Check this https://stackoverflow.com/questions/58084494/golang-how-can-i-get-authorization-from-mux
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

func secureHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func loggingMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[LOG] %s %s %s %d %s", r.RemoteAddr, r.RequestURI, r.Method, r.ContentLength, r.Header.Get("User-Agent"))

		h.ServeHTTP(w, r)
	})
}

// setRouters sets the all required routers
func (app *App) setRouters() {

	//Companies
	app.Get("/companies", corsHandler(app.handleRequest(handler.GetCompanies)))
	app.Get("/companies/{companyID}", corsHandler(app.handleRequest(handler.GetCompany)))
	app.Post("/companies", corsHandler(app.handleRequest(handler.CreateCompany)))
	app.Post("/companies/{companyID}", corsHandler(app.handleRequest(handler.UpdateCompany)))
	app.Post("/companies/{companyID}/delete", corsHandler(app.handleRequest(handler.DeleteCompany)))

	//Domains
	app.Get("/companies/{companyID}/domains", corsHandler(app.handleRequest(handler.GetDomainsForCompany)))
	app.Get("/companies/{companyID}/domains/{domainID}", corsHandler(app.handleRequest(handler.GetDomainForCompany)))
	app.Post("/companies/{companyID}/domains", corsHandler(app.handleRequest(handler.CreateDomainForCompany)))
	app.Post("/companies/{companyID}/domains/{domainID}", corsHandler(app.handleRequest(handler.UpdateDomainForCompany)))
	app.Post("/companies/{companyID}/domains/{domainID}/delete", corsHandler(app.handleRequest(handler.DeleteDomainForCompany)))
	app.Get("/domains", corsHandler(app.handleRequest(handler.GetDomains)))
	app.Get("/domains/{domainID}", corsHandler(app.handleRequest(handler.GetDomain)))
	app.Post("/domains", corsHandler(app.handleRequest(handler.CreateDomain)))
	app.Post("/domains/{domainID}", corsHandler(app.handleRequest(handler.UpdateDomain)))
	app.Post("/domains/{domainID}/delete", corsHandler(app.handleRequest(handler.DeleteDomain)))

	//Subdomains
	app.Get("/domains/{domainID}/subdomains", corsHandler(app.handleRequest(handler.GetSubdomainsForDomain)))
	app.Get("/domains/{domainID}/subdomains/{subdomainID}", corsHandler(app.handleRequest(handler.GetSubdomainForDomain)))
	app.Post("/domains/{domainID}/subdomains", corsHandler(app.handleRequest(handler.CreateSubdomainForDomain)))
	app.Post("/domains/{domainID}/subdomains/{subdomainID}", corsHandler(app.handleRequest(handler.UpdateSubdomainForDomain)))
	app.Post("/domains/{domainID}/subdomains/{subdomainID}/delete", corsHandler(app.handleRequest(handler.DeleteSubdomainForDomain)))
	app.Get("/subdomains", corsHandler(app.handleRequest(handler.GetSubdomains)))
	app.Get("/subdomains/{subdomainID}", corsHandler(app.handleRequest(handler.GetSubdomain)))
	app.Post("/subdomains", corsHandler(app.handleRequest(handler.CreateSubdomain)))
	app.Post("/subdomains/{subdomainID}", corsHandler(app.handleRequest(handler.UpdateSubdomain)))
	app.Post("/subdomains/{subdomainID}/delete", corsHandler(app.handleRequest(handler.DeleteSubdomain)))

	//ApiKeys
	app.Get("/apikeys", corsHandler(app.handleRequest(handler.GetApiKeys)))
	app.Get("/apikeys/{apiKeyID}", corsHandler(app.handleRequest(handler.GetApiKey)))
	app.Post("/apikeys/{apiKeyID}", corsHandler(app.handleRequest(handler.UpdateApiKey)))
	app.Post("/apikeys", corsHandler(app.handleRequest(handler.CreateApiKey)))
	app.Post("/apikeys/{apiKeyID}/delete", corsHandler(app.handleRequest(handler.DeleteApiKey)))

	//Users
	app.Get("/users", corsHandler(app.handleRequest(handler.GetUsers)))
	app.Get("/users/{userID}", corsHandler(app.handleRequest(handler.GetUser)))
	app.Post("/users/{userID}", corsHandler(app.handleRequest(handler.UpdateUser)))
	app.Post("/users", corsHandler(app.handleRequest(handler.CreateUser)))
	app.Post("/users/{userID}/delete", corsHandler(app.handleRequest(handler.DeleteUser)))
	app.Post("/login", corsHandler(app.handleRequest(handler.Login)))

	//Dummy Login
	//app.Post("/login", corsHandler(app.handleRequest(handler.DummyLogin)))

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
