package usecases

import (
	"fmt"
	"log/slog"
	"myclothing/app/product/entities"
	"myclothing/app/product/repositories"
	"time"
)

type productUsecaseImpl struct {
	productRepo repositories.ProductRepository
}

func NewProductUsecaseImpl(productRepo repositories.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{productRepo: productRepo}
}

func (u *productUsecaseImpl) GetProducts() ([]entities.Product, error) {
	products, err := u.productRepo.SelectProducts()
	if err != nil {
		slog.Debug(fmt.Sprintf("%v products returned", len(products)))
		return nil, err
	}

	return products, nil
}

func (u *productUsecaseImpl) CreateProduct(product *entities.Product) error {
	u.GenerateSku(product)

	err := u.productRepo.InsertProduct(product)
	if err != nil {
		slog.Debug("Product created with ID:", product.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) CreateProductCategory(productCategory *entities.ProductCategory) error {
	err := u.productRepo.InsertProductCategory(productCategory)
	if err != nil {
		slog.Debug("Product category created with ID:", productCategory.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) CreateProductSubcategory(productSubcategory *entities.ProductSubcategory) error {
	err := u.productRepo.InsertProductSubcategory(productSubcategory)
	if err != nil {
		slog.Debug("Product subcategory created with ID:", productSubcategory.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) CreateProductSource(productSource *entities.ProductSource) error {
	err := u.productRepo.InsertProductSource(productSource)
	if err != nil {
		slog.Debug("Product source created with ID:", productSource.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) GenerateSku(product *entities.Product) {
	timestamp := time.Now().Unix() % 10000
	id := product.Id % 1000000

	sku := fmt.Sprintf("%04d%06d", timestamp, id)
	product.Sku = sku
}
