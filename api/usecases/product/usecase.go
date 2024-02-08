package product

import (
	"myclothing/api/domain/product"
)

type Usecase interface {
	GetProducts() ([]product.Product, error)
	CreateProduct(product *product.Product) error
	CreateProductCategory(productCategory *product.ProductCategory) error
	CreateProductSubcategory(productSubcategory *product.ProductSubcategory) error
	CreateProductSource(productSource *product.ProductSource) error
}
