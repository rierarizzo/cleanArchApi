package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	appError "myclothing/api/domain/error"
	productDomain "myclothing/api/domain/product"
	productRepo "myclothing/api/persistence/product"
	sqlc2 "myclothing/api/persistence/sqlc"
	"strconv"
)

type productPostgresRepository struct {
	db             *sql.DB
	productQueries *sqlc2.Queries
	ctx            context.Context
}

func NewProductPostgresRepository(db *sql.DB) productRepo.Repository {
	productQueries := sqlc2.New(db)
	ctx := context.Background()

	return &productPostgresRepository{
		db:             db,
		productQueries: productQueries,
		ctx:            ctx,
	}
}

func (r *productPostgresRepository) SelectProducts() ([]productDomain.Product, error) {
	products := make([]productDomain.Product, 0)

	productRows, err := r.productQueries.GetProducts(r.ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			slog.Debug("No rows in product schema.")
			return products, nil
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

func (r *productPostgresRepository) SelectProductCategoryById(categoryId int32) (*productDomain.Category, error) {
	row, err := r.productQueries.GetProductCategoryById(r.ctx, categoryId)
	if err != nil {
		slog.Error("SelectProductCategoryById:", err)
		return nil, appError.ErrRepository
	}

	category := &productDomain.Category{
		Id:          int(row.ID),
		Name:        row.Name,
		Description: row.Description,
	}

	return category, nil
}

func (r *productPostgresRepository) SelectProductSubcategoryById(subcategoryId int32) (*productDomain.Subcategory, error) {
	row, err := r.productQueries.GetProductSubcategoryById(r.ctx, subcategoryId)
	if err != nil {
		slog.Error("SelectProductSubcategoryById:", err)
		return nil, appError.ErrRepository
	}

	category, err := r.SelectProductCategoryById(row.ParentCategoryID)
	if err != nil {
		return nil, err
	}

	subcategory := &productDomain.Subcategory{
		Id:             int(row.ID),
		ParentCategory: category,
		Name:           row.Name,
		Description:    row.Description,
	}

	return subcategory, nil
}

func (r *productPostgresRepository) SelectProductSubcategoryByCategoryId(categoryId int32) ([]productDomain.Subcategory, error) {
	rows, err := r.productQueries.GetProductSubcategoryByCategoryId(r.ctx, categoryId)
	if err != nil {
		slog.Error("SelectProductSubcategoryByCategoryId:", err)
		return nil, appError.ErrRepository
	}

	category, err := r.SelectProductCategoryById(categoryId)
	if err != nil {
		return nil, err
	}

	subcategories := make([]productDomain.Subcategory, 0)
	for _, row := range rows {
		subcategory := &productDomain.Subcategory{
			Id:             int(row.ID),
			ParentCategory: category,
			Name:           row.Name,
			Description:    row.Description,
		}
		subcategories = append(subcategories, *subcategory)
	}

	return subcategories, nil
}

func (r *productPostgresRepository) SelectProductSizeByCode(sizeCode string) (*productDomain.Size, error) {
	row, err := r.productQueries.GetProductSizeByCode(r.ctx, sizeCode)
	if err != nil {
		slog.Error("SelectProductSizeByCode:", err)
		return nil, appError.ErrRepository
	}

	size := &productDomain.Size{
		Code:        row.Code,
		Description: row.Description,
	}

	return size, nil
}

func (r *productPostgresRepository) SelectProductColorById(colorId int32) (*productDomain.Color, error) {
	row, err := r.productQueries.GetProductColorById(r.ctx, colorId)
	if err != nil {
		slog.Error("SelectProductColorById:", err)
		return nil, err
	}

	color := &productDomain.Color{
		Id:   int(row.ID),
		Name: row.Name,
		Hex:  row.Hex,
	}

	return color, nil
}

func (r *productPostgresRepository) SelectProductSourceById(sourceId int32) (*productDomain.Source, error) {
	row, err := r.productQueries.GetProductSourceById(r.ctx, sourceId)
	if err != nil {
		slog.Error("SelectProductSourceById:", err)
		return nil, err
	}

	source := &productDomain.Source{
		Id:   int(row.ID),
		Name: row.Name,
	}

	return source, nil
}

func (r *productPostgresRepository) InsertProduct(product *productDomain.Product) error {
	description := sql.NullString{Valid: false}
	if product.Description != nil {
		description.String = *product.Description
		description.Valid = true
	}

	sourceUrl := sql.NullString{Valid: false}
	if product.SourceUrl != nil {
		sourceUrl.String = *product.SourceUrl
		sourceUrl.Valid = true
	}

	offerPercent := sql.NullInt32{Valid: false}
	if product.OfferPercent != nil {
		offerPercent.Int32 = int32(*product.OfferPercent)
		offerPercent.Valid = true
	}

	row, err := r.productQueries.CreateProduct(r.ctx, sqlc2.CreateProductParams{
		SubcategoryID: int32(product.Subcategory.Id),
		Name:          product.Name,
		Description:   description,
		Price:         fmt.Sprintf("%f", product.Price),
		Cost:          fmt.Sprintf("%f", product.Cost),
		Quantity:      int32(product.Quantity),
		SizeCode:      product.Size.Code,
		ColorID:       int32(product.Color.Id),
		Brand:         product.Brand,
		Sku:           product.Sku,
		Upc:           product.Upc,
		ImageUrl:      product.ImageUrl,
		SourceID:      int32(product.Source.Id),
		SourceUrl:     sourceUrl,
		IsOffered:     product.IsOffered,
		OfferPercent:  offerPercent,
	})
	if err != nil {
		slog.Error("InsertProduct:", err)
		return err
	}

	product.Id = int(row.ID)

	return nil
}

func (r *productPostgresRepository) InsertProductCategory(productCategory *productDomain.Category) error {
	row, err := r.productQueries.CreateProductCategory(r.ctx, sqlc2.CreateProductCategoryParams{
		Name:        productCategory.Name,
		Description: productCategory.Description,
	})
	if err != nil {
		slog.Error("InsertProductCategory:", err)
		return err
	}

	productCategory.Id = int(row.ID)

	return nil
}

func (r *productPostgresRepository) InsertProductSubcategory(productSubcategory *productDomain.Subcategory) error {
	row, err := r.productQueries.CreateProductSubcategory(r.ctx, sqlc2.CreateProductSubcategoryParams{
		ParentCategoryID: int32(productSubcategory.ParentCategory.Id),
		Name:             productSubcategory.Name,
		Description:      productSubcategory.Description,
	})
	if err != nil {
		slog.Error("InsertProductSubcategory:", err)
		return err
	}

	productSubcategory.Id = int(row.ID)

	return nil
}

func (r *productPostgresRepository) InsertProductSource(productSource *productDomain.Source) error {
	row, err := r.productQueries.CreateProductSource(r.ctx, productSource.Name)
	if err != nil {
		slog.Error("InsertProductSource:", err)
		return err
	}

	productSource.Id = int(row.ID)

	return nil
}

func (r *productPostgresRepository) rowToUser(row sqlc2.ProductView) (*productDomain.Product, error) {
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

	offerPercent := int(row.ProductOfferPercent.Int32)

	return &productDomain.Product{
		Id: int(row.ProductID),
		Category: productDomain.Category{
			Id:          int(row.CategoryID),
			Name:        row.CategoryName,
			Description: row.CategoryDescription,
		},
		Subcategory: productDomain.Subcategory{
			Id:          int(row.SubcategoryID),
			Name:        row.SubcategoryName,
			Description: row.SubcategoryDescription,
		},
		Name:        row.ProductName,
		Description: &row.ProductDescription.String,
		Price:       price,
		Cost:        cost,
		Quantity:    int(row.ProductQuantity),
		Size: productDomain.Size{
			Code:        row.SizeCode,
			Description: row.SizeDescription,
		},
		Color: productDomain.Color{
			Id:   int(row.ColorID),
			Name: row.ColorName,
			Hex:  row.ColorHex,
		},
		Brand:    row.ProductBrand,
		Sku:      row.ProductSku,
		Upc:      row.ProductUpc,
		ImageUrl: row.ProductImageUrl,
		Source: productDomain.Source{
			Id:   int(row.SourceID),
			Name: row.SourceName,
		},
		SourceUrl:    &row.ProductSourceUrl.String,
		IsOffered:    row.ProductIsOffered,
		OfferPercent: &offerPercent,
		IsActive:     row.ProductIsActive,
		CreatedAt:    row.ProductCreatedAt.Time,
		UpdatedAt:    row.ProductUpdatedAt.Time,
	}, nil
}
