package main

import "github.com/ikluayping/finalexam/customer"

func main() {
	router := customer.SetupRouter()
	router.Run(":2019")
}
