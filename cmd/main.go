package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	companyHandler "xm-challenge/internal/http"

	"github.com/gorilla/mux"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token, Authorization, Client-id, Client-secret, Total, total")
	headers.Add("Access-Control-Allow-Methods", "GET, PUT, DELETE, POST, OPTIONS, PATCH")
	json.NewEncoder(w)
}

func main() {
	fmt.Println("Starting the application...")

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(optionsHandler)
	// r.HandleFunc("/users", mid.AuthMiddleware(httphandlers.GetUser)).Queries("min", "{min}").Queries("max", "{max}").Queries("email", "{email}").Queries("id", "{id}").Methods("GET")
	// r.HandleFunc("/company", mid.AuthMiddleware(httphandlers.UpdateUser)).Methods("PUT")

	r.HandleFunc("/company", companyHandler.GetCompanyHandler).Methods("GET").Queries("id", "{id}")
	r.HandleFunc("/company", companyHandler.PatchCompanyHandler).Methods("PATCH")
	r.HandleFunc("/company", companyHandler.CreateCompanyHandler).Methods("POST")
	r.HandleFunc("/company", companyHandler.DeleteCompanyHandler).Methods("DELETE")

}
