package server

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	productRepositories "myclothing/api/persistence/product/postgres"
	productHandler "myclothing/api/server/handlers/product/http"
	productUsecases "myclothing/api/usecases/product/default"
	"myclothing/config"
	"net/http"
)

type echoServer struct {
	app *echo.Echo
	mux *http.ServeMux
	db  *sql.DB
	cfg *config.App
}

func NewEchoServer(cfg *config.App, db *sql.DB) Server {
	return &echoServer{
		app: echo.New(),
		mux: http.NewServeMux(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	//s.app.HTTPErrorHandler = echo2.CustomHTTPErrorHandler

	s.initializeProductHttpHandler()

	serverUrl := fmt.Sprintf(":%d", s.cfg.Port)
	//s.app.Logger.Fatal(s.app.Start(serverUrl))
	if err := http.ListenAndServe(serverUrl, s.mux); err != nil {
		panic(err)
	}
}

func (s *echoServer) initializeProductHttpHandler() {
	productPostgresRepo := productRepositories.NewProductPostgresRepository(s.db)
	productUsecase := productUsecases.NewProductUsecaseImpl(productPostgresRepo)

	productHttpHandler := productHandler.NewProductHttpHandler(productUsecase)

	// Routers
	s.mux.HandleFunc("/v1/product", productHttpHandler.GetProducts)
	s.mux.HandleFunc("POST /v1/product", productHttpHandler.CreateProduct)
	s.mux.HandleFunc("POST /v1/product/category", productHttpHandler.CreateProductCategory)
	s.mux.HandleFunc("POST /v1/product/subcategory", productHttpHandler.CreateProductSubcategory)
	s.mux.HandleFunc("POST /v1/product/source", productHttpHandler.CreateProductSource)
}
