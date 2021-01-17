package models

import (
	"time"
)

// Accounts is the account model
type Account struct {
	Owner     string    `json:"owner"`
	Balance   float64   `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

// ErrorResponse define the error json to send
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
