package repositories

import "myclothing/app/user/entities"

type UserRepository interface {
	SelectAppUsersData() ([]entities.AppUser, error)
}
