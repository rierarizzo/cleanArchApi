package main

import (
	"myclothing/config"
	"myclothing/persistence"
	"myclothing/server"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := persistence.NewPostgresDatabase(&cfg.Db)

	server.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
