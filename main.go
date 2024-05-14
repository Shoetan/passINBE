package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shoetan/passIn/controllers"
	"github.com/shoetan/passIn/database"
	"github.com/shoetan/passIn/middleware"
)

func main() {

	//start database connection

	db, err := database.DB()

	if err != nil {
		log.Fatal("Can't connect to database", err.Error())
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/login", controllers.Login(db)).Methods("POST")

	router.HandleFunc("/api/register", controllers.Register(db)).Methods("POST")

	router.HandleFunc("/api/record/{id}", controllers.AddRecord(db)).Methods("POST")

	router.HandleFunc("/api/records", controllers.GetRecords(db)).Methods("GET")
	router.HandleFunc("/api/updateRecord/{id}", controllers.PatchRecord(db)).Methods("PATCH")

	fmt.Println("Server is Running on port 3001")

	log.Fatal(http.ListenAndServe(":3001", middleware.EnableCors(router)))

}
