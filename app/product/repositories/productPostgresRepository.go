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
)

type productPostgresRepository struct {
	productQueries *sqlc.Queries
}

func NewProductPostgresRepository(db *sql.DB) ProductRepository {
	productQueries := sqlc.New(db)

	return &productPostgresRepository{productQueries: productQueries}
}

func (r *productPostgresRepository) SelectProducts() ([]entities.Product, error) {
	products := make([]entities.Product, 0)

	productRows, err := r.productQueries.GetProjects(context.TODO())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return products, nil
		}
		slog.Error("Unknown repository error:", err)
		return nil, appError.ErrRepository
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

func (r *productPostgresRepository) SelectProductCategoryById(categoryId int) (*entities.ProductCategory, error) {
	row, err := r.productQueries.GetProductCategoryById(context.TODO(), int32(categoryId))
	if err != nil {
		slog.Error("Unknown repository error:", err)
		return nil, err
	}

	category := &entities.ProductCategory{
		Id:          int(row.ID),
		Name:        row.Name,
		Description: row.Description,
	}

	// TODO
	// Lógica para traer todos los parents embebidos.

	return category, nil
}

func (r *productPostgresRepository) SelectProductSizeByCode(sizeCode string) (*entities.ProductSize, error) {
	row, err := r.productQueries.GetProductSizeByCode(context.TODO(), sizeCode)
	if err != nil {
		slog.Error("Unknown repository error:", err)
		return nil, err
	}

	size := &entities.ProductSize{
		Code:        row.Code,
		Description: row.Description,
	}

	return size, nil
}

func (r *productPostgresRepository) SelectProductColorById(colorId int) (*entities.ProductColor, error) {
	row, err := r.productQueries.GetProductColorById(context.TODO(), int32(colorId))
	if err != nil {
		slog.Error("Unknown repository error:", err)
		return nil, err
	}

	color := &entities.ProductColor{
		Id:   int(row.ID),
		Name: row.Name,
		Hex:  row.Hex,
	}

	return color, nil
}

func (r *productPostgresRepository) SelectProductSourceById(sourceId int) (*entities.ProductSource, error) {
	row, err := r.productQueries.GetProductSourceById(context.TODO(), int32(sourceId))
	if err != nil {
		slog.Error("Unknown repository error:", err)
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
		Id: int(row.ID),
		Category: entities.ProductCategory{
			Id: int(row.CategoryID),
		},
		Name:        row.Name,
		Description: row.Description.String,
		Price:       price,
		Cost:        cost,
		Size: entities.ProductSize{
			Code: row.SizeCode,
		},
		Color: entities.ProductColor{
			Id: int(row.ColorID),
		},
		Brand:    row.Brand,
		Sku:      row.Sku,
		Upc:      row.Upc,
		ImageUrl: row.ImageUrl,
		Source: entities.ProductSource{
			Id: int(row.SourceID),
		},
		SourceUrl:    row.SourceUrl.String,
		Offer:        row.Offer,
		OfferPercent: int(row.OfferPercent.Int32),
		Active:       row.Active,
		CreatedAt:    row.CreatedAt.Time,
		UpdatedAt:    row.UpdatedAt.Time,
	}, nil
}
