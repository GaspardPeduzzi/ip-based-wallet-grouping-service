package tests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	api "../api"
	users "../users"
	utils "../utils"
)

var i []interface{}

func TestUserHash(t *testing.T) {
	user1 := users.User{IPSRC: "95.40.123.251", WalletID: "4e844548-a23c-4a9d-8d4f-f137cd93c88c", Payload: i}
	user2 := users.User{IPSRC: "95.40.123.251", WalletID: "79ce17c9-97bf-417d-a181-54b9646bb79a", Payload: i}
	user3 := users.User{IPSRC: "221.29.221.1301", WalletID: "861886ef-2d04-4145-a503-277c1ecd0773", Payload: i}
	t.Run("User 1", testUserHashfunc(user1, "f66e5d92198b0173e216f7383cf32c2aeba7263b732995c168b4a1eceb833f13"))
	t.Run("User 2", testUserHashfunc(user2, "f66e5d92198b0173e216f7383cf32c2aeba7263b732995c168b4a1eceb833f13"))
	t.Run("User 3", testUserHashfunc(user3, "9c56e49a66f8fb6aaa50d0aaaff74d43ae123035f4878f45635fb94e1c1b4207"))

}

func testUserHashfunc(user users.User, expectedHash string) func(*testing.T) {
	return func(t *testing.T) {
		hash := utils.GetUserHash(user)
		if expectedHash != utils.GetUserHash(user) {
			t.Errorf("Invalid computed hash:" + "Obtained:" + hash + "Expected" + expectedHash)
		}
	}
}
func TestIsUserStoredInDB(t *testing.T) {

	expectedDB := map[string][]string{
		"f66e5d92198b0173e216f7383cf32c2aeba7263b732995c168b4a1eceb833f13": {"4e844548-a23c-4a9d-8d4f-f137cd93c88c", "79ce17c9-97bf-417d-a181-54b9646bb79a"},
		"9c56e49a66f8fb6aaa50d0aaaff74d43ae123035f4878f45635fb94e1c1b4207": {"861886ef-2d04-4145-a503-277c1ecd0773"},
	}

	user1 := users.User{IPSRC: "95.40.123.251", WalletID: "4e844548-a23c-4a9d-8d4f-f137cd93c88c", Payload: i}
	user2 := users.User{IPSRC: "95.40.123.251", WalletID: "79ce17c9-97bf-417d-a181-54b9646bb79a", Payload: i}
	user3 := users.User{IPSRC: "221.29.221.1301", WalletID: "861886ef-2d04-4145-a503-277c1ecd0773", Payload: i}

	api.InitDB()
	utils.StoreUser(user1)
	utils.StoreUser(user2)
	utils.StoreUser(user3)

	if !reflect.DeepEqual(expectedDB, users.UserDB) {
		t.Errorf("Storing the users in memory doesnt work properly")
	}

}

func testIsUserCorreclyClassifiedFunc(users []users.User, expectedReputableUsers []string) func(*testing.T) {
	return func(t *testing.T) {
		api.InitDB()
		for _, user := range users {
			utils.StoreUser(user)
		}
		reputableUsers := utils.GetReputableUsers()

		if len(reputableUsers) == 0 {
			reputableUsers = append(reputableUsers, "")
		}

		if !(reflect.DeepEqual(reputableUsers, expectedReputableUsers)) {
			t.Errorf("Classified users and expected classified users do not match")
		}
	}
}

func TestClassifciationWithFile(t *testing.T) {

	// Get the test file
	testFile, err := os.Open("req_data.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Test file successfully opened")
	defer testFile.Close()
	// Read the opened file as byte array and unmarchal it into our structs
	testsByteValue, _ := ioutil.ReadAll(testFile)
	var testBatch TestsBatch
	json.Unmarshal(testsByteValue, &testBatch)

	// Run each test
	for name, test := range testBatch {
		t.Run(name, testIsUserCorreclyClassifiedFunc(test.Users, test.NonReputableWallets))
	}

}
