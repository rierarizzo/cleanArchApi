package repositories

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	appError "myclothing/app/error"
	"myclothing/app/product/entities"
	"myclothing/database/postgres/sqlc"
	"strconv"
	"time"
)

type productPostgresRepository struct {
	productQueries *sqlc.Queries
	ctxTimeout     context.Context
	cancelFunc     context.CancelFunc
}

func NewProductPostgresRepository(db *sql.DB) ProductRepository {
	productQueries := sqlc.New(db)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*3)

	return &productPostgresRepository{
		productQueries: productQueries,
		ctxTimeout:     ctxTimeout,
		cancelFunc:     cancel,
	}
}

func (r *productPostgresRepository) SelectProducts() ([]entities.Product, error) {
	defer r.cancelFunc()

	products := make([]entities.Product, 0)

	productRows, err := r.productQueries.GetProjects(r.ctxTimeout)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return products, nil
		}
		slog.Error("Unknown repository error:", err)
		return nil, appError.ErrRepository
	}

	for _, row := range productRows {
		categoryChan := make(chan *entities.ProductCategory)
		sizeChan := make(chan *entities.ProductSize)
		colorChan := make(chan *entities.ProductColor)
		sourceChan := make(chan *entities.ProductSource)
		errChan := make(chan error, 4)

		go func() {
			category, err := r.SelectProductCategoryById(row.CategoryID)
			if err != nil {
				errChan <- fmt.Errorf("failed to get product category: %v", err)
				return
			}
			categoryChan <- category
		}()

		go func() {
			size, err := r.SelectProductSizeByCode(row.SizeCode)
			if err != nil {
				errChan <- fmt.Errorf("failed to get product size: %v", err)
				return
			}
			sizeChan <- size
		}()

		go func() {
			color, err := r.SelectProductColorById(row.ColorID)
			if err != nil {
				errChan <- fmt.Errorf("failed to get product color: %v", err)
				return
			}
			colorChan <- color
		}()

		go func() {
			source, err := r.SelectProductSourceById(row.SourceID)
			if err != nil {
				errChan <- fmt.Errorf("failed to get product source: %v", err)
				return
			}
			sourceChan <- source
		}()

		// Esperamos a que todas las goroutines finalicen
		var categoryResult *entities.ProductCategory
		var sizeResult *entities.ProductSize
		var colorResult *entities.ProductColor
		var sourceResult *entities.ProductSource
		for i := 0; i < 4; i++ {
			select {
			case category := <-categoryChan:
				categoryResult = category
			case size := <-sizeChan:
				sizeResult = size
			case color := <-colorChan:
				colorResult = color
			case source := <-sourceChan:
				sourceResult = source
			case err := <-errChan:
				return nil, err
			}
		}

		product, err := r.rowToUser(row)
		if err != nil {
			return nil, err
		}

		product.Category = *categoryResult
		product.Size = *sizeResult
		product.Color = *colorResult
		product.Source = *sourceResult

		products = append(products, *product)
	}

	return products, nil
}

func (r *productPostgresRepository) SelectProductCategoryById(categoryId int32) (*entities.ProductCategory, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductCategoryById(r.ctxTimeout, categoryId)
	if err != nil {
		return nil, err
	}

	category := &entities.ProductCategory{
		Id:             int(row.ID),
		Name:           row.Name,
		ParentCategory: nil,
		Description:    row.Description,
	}

	// There are only 2 levels of categories.
	if row.ParentCategoryID.Valid {
		row, err = r.productQueries.GetProductCategoryById(r.ctxTimeout, row.ParentCategoryID.Int32)
		if err != nil {
			return nil, err
		}

		category.ParentCategory = &entities.ProductCategory{
			Id:             int(row.ID),
			Name:           row.Name,
			ParentCategory: nil,
			Description:    row.Description,
		}
	}

	return category, nil
}

func (r *productPostgresRepository) SelectProductSizeByCode(sizeCode string) (*entities.ProductSize, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductSizeByCode(r.ctxTimeout, sizeCode)
	if err != nil {
		return nil, err
	}

	size := &entities.ProductSize{
		Code:        row.Code,
		Description: row.Description,
	}

	return size, nil
}

func (r *productPostgresRepository) SelectProductColorById(colorId int32) (*entities.ProductColor, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductColorById(r.ctxTimeout, colorId)
	if err != nil {
		return nil, err
	}

	color := &entities.ProductColor{
		Id:   int(row.ID),
		Name: row.Name,
		Hex:  row.Hex,
	}

	return color, nil
}

func (r *productPostgresRepository) SelectProductSourceById(sourceId int32) (*entities.ProductSource, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductSourceById(r.ctxTimeout, sourceId)
	if err != nil {
		return nil, err
	}

	source := &entities.ProductSource{
		Id:   int(row.ID),
		Name: row.Name,
	}

	return source, nil
}

func (r *productPostgresRepository) rowToUser(row sqlc.Product) (*entities.Product, error) {
	price, err := strconv.ParseFloat(row.Price, 64)
	if err != nil {
		slog.Error("Cannot convert price to float64:", err)
		return nil, appError.ErrConversion
	}

	cost, err := strconv.ParseFloat(row.Cost, 64)
	if err != nil {
		slog.Error("Cannot convert cost to float64:", err)
		return nil, appError.ErrConversion
	}

	return &entities.Product{
		Id:           int(row.ID),
		Name:         row.Name,
		Description:  row.Description.String,
		Price:        price,
		Cost:         cost,
		Brand:        row.Brand,
		Sku:          row.Sku,
		Upc:          row.Upc,
		ImageUrl:     row.ImageUrl,
		SourceUrl:    row.SourceUrl.String,
		Offer:        row.Offer,
		OfferPercent: int(row.OfferPercent.Int32),
		Active:       row.Active,
		CreatedAt:    row.CreatedAt.Time,
		UpdatedAt:    row.UpdatedAt.Time,
	}, nil
}
