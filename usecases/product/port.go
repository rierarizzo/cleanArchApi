package product

import (
	"myclothing/entities/product"
)

type Usecase interface {
	GetProducts() ([]product.Product, error)
	CreateProduct(product *product.Product) error
	CreateProductCategory(productCategory *product.Category) error
	CreateProductSubcategory(productSubcategory *product.Subcategory) error
	CreateProductSource(productSource *product.Source) error
}
