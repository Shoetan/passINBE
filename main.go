package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shoetan/passIn/controllers"
	"github.com/shoetan/passIn/dbConfig"
	"github.com/shoetan/passIn/middleware"
	"github.com/gorilla/mux"
)

func main() {

	//start database connection

	db, err := dbconfig.DB()

	if err != nil {
		log.Fatal("Can't connect to database", err.Error())
	}

	router := mux.NewRouter()

	router.HandleFunc("/api/login", controllers.Login(db)).Methods("POST")

	router.HandleFunc("/api/register", controllers.Register(db)).Methods("POST")

	router.HandleFunc("/api/record/{id}", controllers.AddRecord(db)).Methods("POST")

	router.HandleFunc("/api/records", controllers.GetRecords(db)).Methods("GET")



	fmt.Println("Server is Running on port 3001")

	log.Fatal(http.ListenAndServe(":3001", middleware.EnableCors(router)))

}