package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/shoetan/passIn/types"
	"github.com/shoetan/passIn/utils"
)

func Login(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var loginPayload types.Login
		if err := json.NewDecoder(r.Body).Decode(&loginPayload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		row := db.QueryRow("SELECT email, master_password, user_id FROM users WHERE email = $1", loginPayload.Email)
		var email, masterPassword string
		var userId int
		if err := row.Scan(&email, &masterPassword, &userId); err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if err := utils.ComparePassword(masterPassword, loginPayload.Password); err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
			return
		}

		token, tokenErr := utils.CreateJwtToken(loginPayload.Email)
		if tokenErr != nil {
			http.Error(w, tokenErr.Error(), http.StatusInternalServerError)
			return
		}

		response := types.LoginResponse{
			UserID: userId,
			Email: email,
			Token: token,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}
