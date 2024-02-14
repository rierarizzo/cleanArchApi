package product

import (
	"myclothing/domain/product"
)

type Repository interface {
	SelectProducts() ([]product.Product, error)
	SelectProductCategoryById(categoryId int32) (*product.Category, error)
	SelectProductSubcategoryById(subcategoryId int32) (*product.Subcategory, error)
	SelectProductSubcategoryByCategoryId(categoryId int32) ([]product.Subcategory, error)
	SelectProductSizeByCode(sizeCode string) (*product.Size, error)
	SelectProductColorById(colorId int32) (*product.Color, error)
	SelectProductSourceById(sourceId int32) (*product.Source, error)
	InsertProduct(product *product.Product) error
	InsertProductCategory(productCategory *product.Category) error
	InsertProductSubcategory(productSubcategory *product.Subcategory) error
	InsertProductSource(productSource *product.Source) error
}
