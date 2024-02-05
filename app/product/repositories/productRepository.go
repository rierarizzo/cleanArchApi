package repositories

import "myclothing/app/product/entities"

type ProductRepository interface {
	SelectProducts() ([]entities.Product, error)
}
