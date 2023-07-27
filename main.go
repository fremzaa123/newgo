package main

import (
	"fmt"
	"newgo-app/connect"

	_ "github.com/lib/pq"
)

func main() {
	db, err := connect.ConnectDB()
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

}
