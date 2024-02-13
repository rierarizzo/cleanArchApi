package product

import (
	"fmt"
	"log/slog"
	productDomain "myclothing/entities/product"
	productRepo "myclothing/persistence/product"
	"time"
)

type productUsecaseImpl struct {
	productRepo productRepo.Repository
}

func NewProductUsecaseImpl(productRepo productRepo.Repository) Usecase {
	return &productUsecaseImpl{productRepo: productRepo}
}

func (u *productUsecaseImpl) GetProducts() ([]productDomain.Product, error) {
	products, err := u.productRepo.SelectProducts()
	if err != nil {
		slog.Debug(fmt.Sprintf("%v products returned", len(products)))
		return nil, err
	}

	return products, nil
}

func (u *productUsecaseImpl) CreateProduct(product *productDomain.Product) error {
	u.GenerateSku(product)

	err := u.productRepo.InsertProduct(product)
	if err != nil {
		slog.Debug("Product created with ID:", product.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) CreateProductCategory(productCategory *productDomain.Category) error {
	err := u.productRepo.InsertProductCategory(productCategory)
	if err != nil {
		slog.Debug("Product category created with ID:", productCategory.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) CreateProductSubcategory(productSubcategory *productDomain.Subcategory) error {
	err := u.productRepo.InsertProductSubcategory(productSubcategory)
	if err != nil {
		slog.Debug("Product subcategory created with ID:", productSubcategory.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) CreateProductSource(productSource *productDomain.Source) error {
	err := u.productRepo.InsertProductSource(productSource)
	if err != nil {
		slog.Debug("Product source created with ID:", productSource.Id)
		return err
	}

	return nil
}

func (u *productUsecaseImpl) GenerateSku(product *productDomain.Product) {
	timestamp := time.Now().Unix() % 10000
	id := product.Id % 1000000

	sku := fmt.Sprintf("%04d%06d", timestamp, id)
	product.Sku = sku
}
