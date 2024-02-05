package usecases

import "myclothing/app/product/entities"

type ProductUsecase interface {
	GetProducts() ([]entities.Product, error)
}
