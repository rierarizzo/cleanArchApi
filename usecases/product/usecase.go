package product

import (
	"fmt"
	"log/slog"
	productDomain "myclothing/domain/product"
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
		return nil, err
	}

	slog.Debug(fmt.Sprintf("%v products returned", len(products)))
	return products, nil
}

func (u *productUsecaseImpl) CreateProduct(product *productDomain.Product) error {
	u.GenerateSku(product)

	err := u.productRepo.InsertProduct(product)
	if err != nil {
		return err
	}

	slog.Debug("Product created with ID:", product.Id)
	return nil
}

func (u *productUsecaseImpl) CreateProductCategory(productCategory *productDomain.Category) error {
	err := u.productRepo.InsertProductCategory(productCategory)
	if err != nil {
		return err
	}

	slog.Debug("Product category created with ID:", productCategory.Id)
	return nil
}

func (u *productUsecaseImpl) CreateProductSubcategory(productSubcategory *productDomain.Subcategory) error {
	err := u.productRepo.InsertProductSubcategory(productSubcategory)
	if err != nil {
		return err
	}

	slog.Debug("Product subcategory created with ID:", productSubcategory.Id)
	return nil
}

func (u *productUsecaseImpl) CreateProductColor(productColor *productDomain.Color) error {
	err := u.productRepo.InsertProductColor(productColor)
	if err != nil {
		return err
	}

	slog.Debug("Product color created with ID:", productColor.Id)
	return nil
}

func (u *productUsecaseImpl) CreateProductSource(productSource *productDomain.Source) error {
	err := u.productRepo.InsertProductSource(productSource)
	if err != nil {
		return err
	}

	slog.Debug("Product source created with ID:", productSource.Id)
	return nil
}

func (u *productUsecaseImpl) GenerateSku(product *productDomain.Product) {
	timestamp := time.Now().Unix() % 10000
	id := product.Id % 1000000

	sku := fmt.Sprintf("%04d%06d", timestamp, id)
	product.Sku = sku
}
