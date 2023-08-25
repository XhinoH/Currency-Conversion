package main

import (
	"gocapri/db"
	"gocapri/routes"
)

func main() {
	db.Init()
	r := routes.Init()
	r.Run("localhost:8080")
}
