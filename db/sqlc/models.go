// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"time"
)

type Accounts struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
}

type Entries struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfers struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// it must be positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}