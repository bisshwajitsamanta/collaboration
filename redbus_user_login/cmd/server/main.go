package main

import (
	"fmt"
	"redbus_user_login/internal/db"
)

func Run() error {
	fmt.Println("Redbus Ticketing Service")
	_, err := db.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to Database")
		return err
	}
	fmt.Println("Successfully connected to Database")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Unable to start Redbus Application")
		fmt.Println(err)
	}
}
