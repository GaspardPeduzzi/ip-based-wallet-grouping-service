package tests

import (
	users "../users"
)

// TestUser : Represent a test with a list of users
type TestUsers struct {
	NonReputableWallets []string     `json:"non_reputable_wallets"`
	Users               []users.User `json:"data"`
}

// TestsBatch : Represent a serie of tests
type TestsBatch map[string]TestUsers
