package usecases

import "myclothing/app/product/entities"

type ProductUsecase interface {
	GetProducts() ([]entities.Product, error)
	CreateProduct(product entities.Product) (*entities.Product, error)
	CreateProductCategory(productCategory entities.ProductCategory) (*entities.ProductCategory, error)
	CreateProductSubcategory(productSubcategory entities.ProductSubcategory) (*entities.ProductSubcategory, error)
	CreateProductSource(productSource entities.ProductSource) (*entities.ProductSource, error)
}
