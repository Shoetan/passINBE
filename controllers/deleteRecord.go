package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shoetan/passIn/utils"
)

func DeleteRecord(db *sql.DB)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == ""{
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return 
		}

		tokenString = tokenString[len("Bearer "):]

		err := utils.VerifyJwtToken(tokenString)

		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		id:=mux.Vars(r)["id"]

		recordId, _ := strconv.Atoi(id)

		_,err = db.Exec("DELETE FROM vault WHERE record_id = $1", recordId)

		if err != nil{
			http.Error(w, "Failed to delete record", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)












	}
}

