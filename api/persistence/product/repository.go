package product

import (
	"myclothing/api/domain/product"
)

type Repository interface {
	SelectProducts() ([]product.Product, error)
	SelectProductCategoryById(categoryId int32) (*product.ProductCategory, error)
	SelectProductSubcategoryById(subcategoryId int32) (*product.ProductSubcategory, error)
	SelectProductSubcategoryByCategoryId(categoryId int32) ([]product.ProductSubcategory, error)
	SelectProductSizeByCode(sizeCode string) (*product.ProductSize, error)
	SelectProductColorById(colorId int32) (*product.ProductColor, error)
	SelectProductSourceById(sourceId int32) (*product.ProductSource, error)
	InsertProduct(product *product.Product) error
	InsertProductCategory(productCategory *product.ProductCategory) error
	InsertProductSubcategory(productSubcategory *product.ProductSubcategory) error
	InsertProductSource(productSource *product.ProductSource) error
}
