package main

import (
	"go-echo-mysql/db"
	"go-echo-mysql/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
