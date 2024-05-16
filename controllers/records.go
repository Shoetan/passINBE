package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"

	"github.com/shoetan/passIn/utils"
)


type RecordDetails struct {
	ID uint `json:"id"`
	UserId uint `json:"user_id"`
	RecordName string `json:"record_name"`
	RecordPassword string `json:"record_password"`
	RecordEmail string `json:"record_email"`
}



func GetRecords(db * sql.DB) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){

		tokenString := r.Header.Get("Authorization")

		if tokenString == ""{
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]


		err := utils.VerifyJwtToken(tokenString)

		if err != nil{
			http.Error(w,"Invalid tokenn", http.StatusUnauthorized)
			return
		}

		id := mux.Vars(r)["id"]
		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}

		var keyphrase = []byte(utils.GetSecretKey("KEY_PHRASE"))

		rows, err := db.Query("SELECT * FROM vault WHERE user_id = $1", userID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer rows.Close()


		records := []RecordDetails{}

		for rows.Next(){
			var r RecordDetails

			if err := rows.Scan(&r.ID, &r.UserId, &r.RecordName, &r.RecordPassword, &r.RecordEmail); err != nil {
				fmt.Println(err)
			}

			//decryptedPwd := utils.DecryptPassword([]byte(r.RecordPassword),  string(keyphrase))

			decry := string(utils.DecryptPassword([]byte (r.RecordPassword), string(keyphrase)))

			recordResponse := RecordDetails{
				ID : r.ID,
				UserId: r.UserId,
				RecordName: r.RecordName,
				RecordPassword: decry,
				RecordEmail: r.RecordEmail,
			}

			records = append(records, recordResponse)

		}

		err = json.NewEncoder(w).Encode(records)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}