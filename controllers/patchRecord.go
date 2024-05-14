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

type updateResponse struct {
	RecordId int `json:"record_id"`
	RecordName string `json:"record_name"`
	RecordEmail string `json:"record_emai"`
}


func PatchRecord(db *sql.DB)http.HandlerFunc{
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

		var keyphrase = []byte(utils.GetSecretKey("KEY_PHRASE"))

		encryptedPwd := utils.EncryptPassword([]byte(recordPayload.RecordPassword), string(keyphrase))

		id:=mux.Vars(r)["id"]

		recordId, _ := strconv.Atoi(id)

		//check if record Id exists

		var existingRecordId int

		err = db.QueryRow("SELECT record_id FROM vault WHERE record_id = $1", recordId).Scan(&existingRecordId)

		switch{
		case err == sql.ErrNoRows:
			http.Error(w, "Record does not exist", http.StatusBadRequest)
			return
		case err != nil:
			http.Error(w, err.Error(),http.StatusInternalServerError)
			return
		}

		_, err = db.Exec("UPDATE vault SET record_name =$1, record_password=$2, record_email=$3 WHERE record_id = $4",recordPayload.RecordName, encryptedPwd, recordPayload.RecordEmail, recordId  )

		if err != nil{
			http.Error(w, "Failed to update record", http.StatusInternalServerError)
			return
		}

		updateRes := updateResponse{
			RecordId: existingRecordId ,
			RecordName: recordPayload.RecordName,
			RecordEmail: recordPayload.RecordEmail,
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updateRes)

	}

}