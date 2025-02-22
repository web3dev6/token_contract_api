// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type Token struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Address   string `json:"address"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Amount    string `json:"amount"`
	Owner     string `json:"owner"`
	Authority string `json:"authority"`
}

type Transaction struct {
	ID          int64           `json:"id"`
	Username    string          `json:"username"`
	Context     string          `json:"context"`
	Payload     json.RawMessage `json:"payload"`
	IsConfirmed bool            `json:"is_confirmed"`
	Status      string          `json:"status"`
	CreatedAt   time.Time       `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	WalletAddress     string    `json:"wallet_address"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
