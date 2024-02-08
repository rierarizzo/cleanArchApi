package main

import (
	"myclothing/api/persistence"
	"myclothing/api/server"
	"myclothing/config"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := persistence.NewPostgresDatabase(&cfg.Db)

	server.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
