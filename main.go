package main

import (
	"cleanArchApi/config"
	"cleanArchApi/database"
	"cleanArchApi/server"
	"fmt"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg.Db)

	fmt.Println("Hello, World")

	server.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
