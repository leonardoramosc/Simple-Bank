package main

import (
	"bank/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/accounts", handlers.AccountsHandler)
	r.HandleFunc("/accounts/{id}", handlers.GetAccountById)

	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
