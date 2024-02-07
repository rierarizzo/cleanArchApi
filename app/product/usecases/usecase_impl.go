package usecases

import (
	"fmt"
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
	return u.productRepo.SelectProducts()
}

func (u *productUsecaseImpl) CreateProduct(product entities.Product) (*entities.Product, error) {
	u.GenerateSku(&product)

	return u.productRepo.InsertProduct(product)
}

func (u *productUsecaseImpl) CreateProductCategory(productCategory entities.ProductCategory) (*entities.ProductCategory, error) {
	return u.productRepo.InsertProductCategory(productCategory)
}

func (u *productUsecaseImpl) CreateProductSubcategory(productSubcategory entities.ProductSubcategory) (*entities.ProductSubcategory, error) {
	return u.productRepo.InsertProductSubcategory(productSubcategory)
}

func (u *productUsecaseImpl) CreateProductSource(productSource entities.ProductSource) (*entities.ProductSource, error) {
	return u.productRepo.InsertProductSource(productSource)
}

func (u *productUsecaseImpl) GenerateSku(product *entities.Product) {
	timestamp := time.Now().Unix() % 10000
	id := product.Id % 1000000

	sku := fmt.Sprintf("%04d%06d", timestamp, id)
	product.Sku = sku
}
