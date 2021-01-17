package handlers

import (
	database "bank/db"
	models "bank/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// AccountsHandler will serve all the accounts
func AccountsHandler(w http.ResponseWriter, r *http.Request) {
	database.GetConnection()
	defer database.DBCon.Close(context.Background())

	q := "select owner, balance, currency, created_at from accounts"

	rows, err := database.DBCon.Query(context.Background(), q)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	accounts := []models.Account{}

	// rows.Next() returns true if there is an actual row
	//(everytime is called, we will get the next row when calling rows.Scan())
	for i := 0; rows.Next(); i++ {
		var acc models.Account

		// Assing the current row to the Account struct
		rows.Scan(&acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)
		accounts = append(accounts, acc)
	}

	// Convert the slice of accounts into JSON format
	response, err := json.Marshal(accounts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Send the response to the client
	w.Write(response)
}

// GetAccountById fetch and send to the client an account by the ID
func GetAccountById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	database.GetConnection()
	defer database.DBCon.Close(context.Background())

	acc := models.Account{}
	q := "select owner, balance, currency, created_at from accounts where id=$1"
	// Execute Query
	err := database.DBCon.QueryRow(context.Background(), q, id).Scan(
		&acc.Owner,
		&acc.Balance,
		&acc.Currency,
		&acc.CreatedAt,
	)

	w.Header().Set("content-type", "application/json")

	if err != nil {
		fmt.Println(err)
		errorResponse := models.ErrorResponse{
			Status:  "fail",
			Message: err.Error(),
		}
		response, _ := json.Marshal(errorResponse)

		if err.Error() != "no rows in result set" {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

		w.Write(response)
		return
	}

	response, err := json.Marshal(acc)

	if err != nil {
		fmt.Println("unable to convert to JSON")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	// Send the response to the client
	w.Write(response)

}
