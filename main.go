package main

import (
	"myclothing/api/persistence"
	"myclothing/api/server/http"
	"myclothing/config"
)

func main() {
	config.Logger()
	cfg := config.GetConfig()

	db := persistence.NewPostgresDatabase(&cfg.Db)

	http.NewEchoServer(&cfg.App, db.GetDb()).Start()
}
