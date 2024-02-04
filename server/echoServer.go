package server

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	userHandlers "myclothing/app/user/handlers"
	userRepositories "myclothing/app/user/repositories"
	userUsecases "myclothing/app/user/usecases"
	"myclothing/config"
	"myclothing/server/echo/middlewares"
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
	s.app.HTTPErrorHandler = middlewares.CustomHTTPErrorHandler

	s.initializeUserHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.cfg.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeUserHttpHandler() {
	userPostgresRepository := userRepositories.NewUserPostgresRepository(s.db)
	userUsecase := userUsecases.NewUserUsecaseImpl(userPostgresRepository)

	userHttpHandler := userHandlers.NewUserHttpHandler(userUsecase)

	// Routers
	userRouters := s.app.Group("v1/user")
	userRouters.GET("", userHttpHandler.GetUsers)
}
