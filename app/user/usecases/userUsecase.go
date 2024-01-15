package usecases

import "cleanArchApi/app/user/entities"

type UserUsecase interface {
	GetAllUsers() ([]entities.AppUser, error)
}
