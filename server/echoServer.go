package server

import (
	"cleanArchApi/config"
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type echoServer struct {
	app *echo.Echo
	db  *sql.DB
	cfg *config.App
}

func NewEchoServer(cfg *config.App, db *sql.DB) Server {
	return &echoServer{
		app: echo.New(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Logger())

	serverUrl := fmt.Sprintf(":%d", s.cfg.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}
