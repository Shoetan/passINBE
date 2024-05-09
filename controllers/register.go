package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shoetan/passIn/types"
	"github.com/shoetan/passIn/utils"
)

type registerResponse struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	Name       string    `json:"name"`
	InsertedAt time.Time `json:"inserted_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func Register(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload types.User
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Check if email already exists
		var existingEmail string
		err := db.QueryRow("SELECT email FROM users WHERE email = $1", payload.Email).Scan(&existingEmail)
		switch {
		case err == sql.ErrNoRows:
			// Email doesn't exist, proceed with registration
		case err != nil:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		default:
			http.Error(w, "Email already exists", http.StatusBadRequest)
			return
		}

		// Hash the master password
		hashedPwd, err := utils.HashPassword(payload.MasterPassword)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		payload.MasterPassword = hashedPwd

		// Insert user into the database
		err = db.QueryRow("INSERT INTO users (name, master_password, email, inserted_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING user_id", payload.Name, payload.MasterPassword, payload.Email, payload.InsertedAt, payload.UpdatedAt).Scan(&payload.UserID)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			return
		}

		response := registerResponse{
			ID:         payload.UserID,
			Name:       payload.Name,
			Email:      payload.Email,
			InsertedAt: payload.InsertedAt,
			UpdatedAt:  payload.UpdatedAt,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
