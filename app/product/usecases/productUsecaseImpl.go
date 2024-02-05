package usecases

import (
	"myclothing/app/product/entities"
	"myclothing/app/product/repositories"
)

type productUsecaseImpl struct {
	productRepo repositories.ProductRepository
}

func NewProductUsecaseImpl(productRepo repositories.ProductRepository) ProductUsecase {
	return &productUsecaseImpl{productRepo: productRepo}
}

func (u *productUsecaseImpl) GetProducts() ([]entities.Product, error) {
	return u.productRepo.SelectProducts()
}
