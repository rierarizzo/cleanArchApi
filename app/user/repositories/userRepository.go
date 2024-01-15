package repositories

import "cleanArchApi/app/user/entities"

type UserRepository interface {
	SelectAppUsersData() ([]entities.AppUser, error)
}
