package repositories

import (
	"context"
	"database/sql"
	"errors"
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
		category, err := r.SelectProductCategoryById(row.CategoryID)
		if err != nil {
			slog.Error("Failed to get product category:", err)
			return nil, appError.ErrRepository
		}

		size, err := r.SelectProductSizeByCode(row.SizeCode)
		if err != nil {
			slog.Error("Failed to get product size:", err)
			return nil, appError.ErrRepository
		}

		color, err := r.SelectProductColorById(row.ColorID)
		if err != nil {
			slog.Error("Failed to get product color:", err)
			return nil, appError.ErrRepository
		}

		source, err := r.SelectProductSourceById(row.SourceID)
		if err != nil {
			slog.Error("Failed to get product source:", err)
			return nil, appError.ErrRepository
		}

		product, err := r.rowToUser(row)
		if err != nil {
			return nil, err
		}

		product.Category = *category
		product.Size = *size
		product.Color = *color
		product.Source = *source

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
