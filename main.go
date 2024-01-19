package main

import (
	"cleanArchApi/config"
	"cleanArchApi/database"
	"cleanArchApi/server"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg.Db)
	
	server.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
