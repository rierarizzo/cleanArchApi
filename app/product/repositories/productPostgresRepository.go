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

const dbTimeout = time.Second * 5

type productPostgresRepository struct {
	db             *sql.DB
	productQueries *sqlc.Queries
	ctxTimeout     context.Context
	cancelFunc     context.CancelFunc
}

func NewProductPostgresRepository(db *sql.DB) ProductRepository {
	productQueries := sqlc.New(db)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), dbTimeout)

	return &productPostgresRepository{
		db:             db,
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
		switch {
		case errors.Is(err, sql.ErrNoRows):
			slog.Debug("No rows in product schema.")
			return products, nil
		case errors.Is(err, context.DeadlineExceeded):
			slog.Error(fmt.Sprintf("Context timeout. Exceeded %v.", dbTimeout))
			return nil, appError.ErrTimeout
		default:
			slog.Error("SelectProducts:", err)
			return nil, appError.ErrRepository
		}
	}

	for _, row := range productRows {
		product, err := r.rowToUser(row)
		if err != nil {
			return nil, err
		}
		products = append(products, *product)
	}

	return products, nil
}

func (r *productPostgresRepository) SelectProductCategoryById(categoryId int32) (*entities.ProductCategory, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductCategoryById(r.ctxTimeout, categoryId)
	if err != nil {
		slog.Error("SelectProductCategoryById:", err)
		return nil, appError.ErrRepository
	}

	category := &entities.ProductCategory{
		Id:          int(row.ID),
		Name:        row.Name,
		Description: row.Description,
	}

	return category, nil
}

func (r *productPostgresRepository) SelectProductSubcategoryById(subcategoryId int32) (*entities.ProductSubcategory, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductSubcategoryById(r.ctxTimeout, subcategoryId)
	if err != nil {
		slog.Error("SelectProductSubcategoryById:", err)
		return nil, appError.ErrRepository
	}

	category, err := r.SelectProductCategoryById(row.ParentCategoryID)
	if err != nil {
		return nil, err
	}

	subcategory := &entities.ProductSubcategory{
		Id:             int(row.ID),
		ParentCategory: category,
		Name:           row.Name,
		Description:    row.Description,
	}

	return subcategory, nil
}

func (r *productPostgresRepository) SelectProductSubcategoryByCategoryId(categoryId int32) ([]entities.ProductSubcategory, error) {
	defer r.cancelFunc()

	rows, err := r.productQueries.GetProductSubcategoryByCategoryId(r.ctxTimeout, categoryId)
	if err != nil {
		slog.Error("SelectProductSubcategoryByCategoryId:", err)
		return nil, appError.ErrRepository
	}

	category, err := r.SelectProductCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	subcategories := make([]entities.ProductSubcategory, 0)
	for _, row := range rows {
		subcategory := &entities.ProductSubcategory{
			Id:             int(row.ID),
			ParentCategory: category,
			Name:           row.Name,
			Description:    row.Description,
		}
		subcategories = append(subcategories, *subcategory)
	}

	return subcategories, nil
}

func (r *productPostgresRepository) SelectProductSizeByCode(sizeCode string) (*entities.ProductSize, error) {
	defer r.cancelFunc()

	row, err := r.productQueries.GetProductSizeByCode(r.ctxTimeout, sizeCode)
	if err != nil {
		slog.Error("SelectProductSizeByCode:", err)
		return nil, appError.ErrRepository
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
		slog.Error("SelectProductColorById:", err)
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
		slog.Error("SelectProductSourceById:", err)
		return nil, err
	}

	source := &entities.ProductSource{
		Id:   int(row.ID),
		Name: row.Name,
	}

	return source, nil
}

func (r *productPostgresRepository) rowToUser(row sqlc.GetProjectsRow) (*entities.Product, error) {
	price, err := strconv.ParseFloat(row.ProductPrice, 64)
	if err != nil {
		slog.Error("Cannot convert price to float64:", err)
		return nil, appError.ErrConversion
	}

	cost, err := strconv.ParseFloat(row.ProductCost, 64)
	if err != nil {
		slog.Error("Cannot convert cost to float64:", err)
		return nil, appError.ErrConversion
	}

	return &entities.Product{
		Id: int(row.ProductID),
		Category: entities.ProductCategory{
			Id:          int(row.CategoryID),
			Name:        row.CategoryName,
			Description: row.CategoryDescription,
		},
		Subcategory: entities.ProductSubcategory{
			Id:          int(row.SubcategoryID),
			Name:        row.SubcategoryName,
			Description: row.SubcategoryDescription,
		},
		Name:        row.ProductName,
		Description: row.ProductDescription.String,
		Price:       price,
		Cost:        cost,
		Quantity:    int(row.ProductQuantity),
		Size: entities.ProductSize{
			Code:        row.SizeCode,
			Description: row.SizeDescription,
		},
		Color: entities.ProductColor{
			Id:   int(row.ColorID),
			Name: row.ColorName,
			Hex:  row.ColorHex,
		},
		Brand:    row.ProductBrand,
		Sku:      row.ProductSku,
		Upc:      row.ProductUpc,
		ImageUrl: row.ProductImageUrl,
		Source: entities.ProductSource{
			Id:   int(row.SourceID),
			Name: row.SourceName,
		},
		SourceUrl:    row.ProductSourceUrl.String,
		IsOffered:    row.ProductIsOffered,
		OfferPercent: int(row.ProductOfferPercent.Int32),
		IsActive:     row.ProductIsActive,
		CreatedAt:    row.ProductCreatedAt.Time,
		UpdatedAt:    row.ProductUpdatedAt.Time,
	}, nil
}
