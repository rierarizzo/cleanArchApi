package usecases

import "myclothing/app/product/entities"

type ProductUsecase interface {
	GetProducts() ([]entities.Product, error)
	CreateProduct(product *entities.Product) error
	CreateProductCategory(productCategory *entities.ProductCategory) error
	CreateProductSubcategory(productSubcategory *entities.ProductSubcategory) error
	CreateProductSource(productSource *entities.ProductSource) error
}
