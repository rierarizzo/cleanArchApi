package main

import (
	"myclothing/config"
	"myclothing/persistence"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := persistence.NewPostgresDatabase(&cfg.Db)

	NewEchoServer(&cfg.App, db.GetDb()).Start()
}
