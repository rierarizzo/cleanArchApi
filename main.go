package main

import (
	"myclothing/config"
	"myclothing/database"
	"myclothing/server"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg.Db)

	server.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
