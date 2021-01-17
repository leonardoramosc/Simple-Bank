package handlers

import (
	database "bank/db"
	models "bank/models"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

	for i := 0; rows.Next(); i++ {
		var acc models.Account

		rows.Scan(&acc.Owner, &acc.Balance, &acc.Currency, &acc.CreatedAt)
		accounts = append(accounts, acc)
	}

	response, err := json.Marshal(accounts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
