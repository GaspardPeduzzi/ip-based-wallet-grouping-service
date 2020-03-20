package users

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
)

// User : Struct for user data
type User struct {
	IPSRC    string        `json:"ip_src"`
	WalletID string        `json:"wallet_id"`
	Payload  []interface{} `json:"payload"`
}

// UserDB : Data structure of the database that keep count of the users
var UserDB map[string][]string

// ReverseDirectory : Map each wallet_id to its hashed ip for complexity issues
var ReverseDirectory map[string]string

// InitDB : Create temporary datastructure to store users
func InitDB() {
	UserDB = make(map[string][]string)
	ReverseDirectory = make(map[string]string)
	log.Println("User database created")

}

// GroupTreshold : Maximum number of users tolerated per IP
const GroupTreshold = 3

// GetUserHash : Return hash of the user IP
func GetUserHash(user User) string {
	hash := sha256.Sum256([]byte(user.IPSRC))
	return hex.EncodeToString(hash[:])
}

// StoreUser : Store user in memory
func StoreUser(user User) bool {
	hashedIP := GetUserHash(user)
	found := contains(UserDB[hashedIP], user.WalletID)
	if !found {
		UserDB[hashedIP] = append(UserDB[hashedIP], user.WalletID)
		ReverseDirectory[user.WalletID] = hashedIP
		log.Println("User " + user.WalletID + " added to the db")
		return true
	}
	return false
}

// IsUserReputable : Return if the user is reputable or not
func IsUserReputable(walletID string) int {

	hashedIP := ReverseDirectory[walletID]
	if len(hashedIP) == 0 {
		return -1 // User does not exist
	}
	groupLen := len(UserDB[hashedIP])

	if groupLen >= GroupTreshold {
		return 0 // User is not reputable
	}
	return 1 // User is reputable
}

// GetReputableUsers : Return all the reputable users of the current database
func GetReputableUsers() []string {
	reputableUser := []string{}

	for k := range UserDB {
		if !(len(UserDB[k]) < GroupTreshold) {
			reputableUser = append(reputableUser, UserDB[k]...)
		}

	}
	return reputableUser
}

func contains(users []string, user string) bool {
	for _, u := range users {
		if u == user {
			return true
		}
	}
	return false
}
