package main

import (
	"github.com/danilosmaciel/api-go-gin/database"
	"github.com/danilosmaciel/api-go-gin/routes"
)

func main() {

	err := database.Connect()

	if err != nil {
		panic(err)
	}

	r := routes.HandleRequests()

	err = r.Run(":8080")

	if err != nil {
		panic(err)
	}
}
