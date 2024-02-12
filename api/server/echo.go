package server

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	productRepositories "myclothing/api/persistence/product/postgres"
	productHandler "myclothing/api/server/handlers/product/http"
	echo2 "myclothing/api/server/middlewares"
	productUsecases "myclothing/api/usecases/product/default"
	"myclothing/config"
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
	s.app.HTTPErrorHandler = echo2.CustomHTTPErrorHandler

	s.initializeProductHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.cfg.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeProductHttpHandler() {
	productPostgresRepo := productRepositories.NewProductPostgresRepository(s.db)
	productUsecase := productUsecases.NewProductUsecaseImpl(productPostgresRepo)

	productHttpHandler := productHandler.NewProductHttpHandler(productUsecase)

	// Routers
	productRouters := s.app.Group("v1/product")
	productRouters.GET("", productHttpHandler.GetProducts)
	productRouters.POST("", productHttpHandler.CreateProduct)
	productRouters.POST("category", productHttpHandler.CreateProductCategory)
	productRouters.POST("subcategory", productHttpHandler.CreateProductSubcategory)
	productRouters.POST("source", productHttpHandler.CreateProductSource)
}
