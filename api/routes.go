package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	users "../users"
	"github.com/gorilla/mux"
)

// AddUser : Adding a new user to the user database
func AddUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newUser users.User
	err := json.Unmarshal(reqBody, &newUser)
	if err != nil {
		log.Println(err)
	}

	success := users.StoreUser(newUser)

	w.WriteHeader(http.StatusCreated)

	if success {
		w.Write([]byte(`{"user_added": true}`))
	} else {
		w.Write([]byte(`{"user_added": false}`))
	}

}

// checkIfReputable : Check if the user is part of a pool of user (larger than CW=3)
func checkIfReputable(w http.ResponseWriter, r *http.Request) {

	walletID := mux.Vars(r)["wallet_id"]

	isReputable := users.IsUserReputable(walletID)

	w.WriteHeader(http.StatusCreated)
	if isReputable == 1 {
		w.Write([]byte(`{"isReputable": true}`))
	} else if isReputable == 0 {
		w.Write([]byte(`{"isReputable": false}`))
	} else {
		w.Write([]byte(`{"error": invalid user}`))
	}

}
