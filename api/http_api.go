package api

import (
	"log"
	"net/http"

	users "../users"
	"github.com/gorilla/mux"
)

// StartServer : Start the HTTP API
func StartServer() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/submitUserData", AddUser).Methods("POST")
	router.HandleFunc("/isUserReputable/{wallet_id}", checkIfReputable).Methods("GET")

	log.Println("API Started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// InitDB : Create temporary datastructure to store users
func InitDB() {
	users.UserDB = make(map[string][]string)
	users.ReverseDirectory = make(map[string]string)
	log.Println("User database created")

}
