package main

import "fmt"

func Run() error {
	fmt.Println("Redbus Ticketing Service")
	return nil
}

func main() {
	if err := Run(); err != nil {
		fmt.Println("Unable to start Redbus Application")
		fmt.Println(err)
	}
}
