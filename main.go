package main

import (
	api "./api"
	users "./users"
)

func main() {
	users.InitDB()
	api.StartServer()
}
