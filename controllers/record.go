package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/shoetan/passIn/types"
	"github.com/shoetan/passIn/utils"
)

type recordResponse struct {
	UserId     int    `json:"user_id"`
	RecordName string `json:"record"`
}

func AddRecord(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		err := utils.VerifyJwtToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		var recordPayload types.AddRecordPayload
		err = json.NewDecoder(r.Body).Decode(&recordPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}



		id := mux.Vars(r)["id"]
		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		// Check if the user exists
		var existingUserID int
		err = db.QueryRow("SELECT user_id FROM vault WHERE user_id = $1", userID).Scan(&existingUserID)
		switch {
		case err == sql.ErrNoRows:
			http.Error(w, "User does not exist", http.StatusBadRequest)
			return
		case err != nil:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var keyphrase = []byte(utils.GetSecretKey("KEY_PHRASE"))

		encryptedPwd := utils.EncryptPassword([]byte(recordPayload.RecordPassword), string(keyphrase))

		

		// Insert record into the database
		_, err = db.Exec("INSERT INTO vault (user_id, record_name, record_password) VALUES ($1, $2, $3)", userID, recordPayload.RecordName, encryptedPwd)
		if err != nil {
			http.Error(w, "Failed to insert record", http.StatusInternalServerError)
			return
		}

		recordResp := recordResponse{
			UserId:     userID,
			RecordName: recordPayload.RecordName,
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(recordResp)
	}
}
