package repositories

import "myclothing/app/product/entities"

type ProductRepository interface {
	SelectProducts() ([]entities.Product, error)
	SelectProductCategoryById(categoryId int32) (*entities.ProductCategory, error)
	SelectProductSubcategoryById(subcategoryId int32) (*entities.ProductSubcategory, error)
	SelectProductSubcategoryByCategoryId(categoryId int32) ([]entities.ProductSubcategory, error)
	SelectProductSizeByCode(sizeCode string) (*entities.ProductSize, error)
	SelectProductColorById(colorId int32) (*entities.ProductColor, error)
	SelectProductSourceById(sourceId int32) (*entities.ProductSource, error)
}
