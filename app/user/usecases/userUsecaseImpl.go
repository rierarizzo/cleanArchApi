package usecases

import (
	"cleanArchApi/app/user/entities"
	"cleanArchApi/app/user/repositories"
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
