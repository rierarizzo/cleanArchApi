package usecases

import "myclothing/app/user/entities"

type UserUsecase interface {
	GetAllUsers() ([]entities.AppUser, error)
}
