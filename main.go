package main

import (
	"cleanArchApi/config"
	"cleanArchApi/database"
	"cleanArchApi/server"
)

func main() {
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg.Db)

	config.Logger()

	server.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
