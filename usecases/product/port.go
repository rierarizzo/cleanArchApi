package product

import (
	"myclothing/domain/product"
)

type Usecase interface {
	GetProducts() ([]product.Product, error)
	CreateProduct(product *product.Product) error
	CreateProductCategory(productCategory *product.Category) error
	CreateProductSubcategory(productSubcategory *product.Subcategory) error
	CreateProductColor(productColor *product.Color) error
	CreateProductSource(productSource *product.Source) error
}
