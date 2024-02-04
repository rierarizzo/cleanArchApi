package usecases

import (
	"myclothing/app/user/entities"
	"myclothing/app/user/repositories"
)

type userUsecaseImpl struct {
	userRepository repositories.UserRepository
}

func NewUserUsecaseImpl(userRepository repositories.UserRepository) UserUsecase {
	return &userUsecaseImpl{userRepository: userRepository}
}

func (u *userUsecaseImpl) GetAllUsers() ([]entities.AppUser, error) {
	return u.userRepository.SelectAppUsersData()
}
