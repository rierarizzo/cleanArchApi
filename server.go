package main

import (
	"database/sql"
	"fmt"
	"myclothing/config"
	productHandler "myclothing/http/handlers/product"
	productRepositories "myclothing/persistence/product"
	productUsecases "myclothing/usecases/product"
	"net/http"
)

type Server interface {
	Start()
}

type echoServer struct {
	mux *http.ServeMux
	db  *sql.DB
	cfg *config.App
}

func NewEchoServer(cfg *config.App, db *sql.DB) Server {
	return &echoServer{
		mux: http.NewServeMux(),
		db:  db,
		cfg: cfg,
	}
}

func (s *echoServer) Start() {
	s.initializeProductHttpHandler()

	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.cfg.Port), s.mux); err != nil {
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
