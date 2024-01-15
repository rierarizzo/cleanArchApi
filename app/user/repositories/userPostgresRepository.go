package repositories

import (
	"cleanArchApi/app/user/entities"
	"cleanArchApi/database/postgres/sqlc"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type userPostgresRepository struct {
	db *sql.DB
}

func NewUserPostgresRepository(db *sql.DB) UserRepository {
	return &userPostgresRepository{db: db}
}

func (r *userPostgresRepository) SelectAppUsersData() ([]entities.AppUser, error) {
	var appUsers []entities.AppUser
	queries := sqlc.New(r.db)
	appUsersModel, err := queries.GetUsers(context.TODO())

	if err != nil {
		slog.Error(fmt.Sprintf("SelectAppUsersData: %v", err))
		return appUsers, err
	}

	for _, v := range appUsersModel {
		appUsers = append(
			appUsers, entities.AppUser{
				Id:           uint32(v.AppUserID),
				Username:     v.Username,
				Email:        v.Email,
				PasswordHash: v.PasswordHash,
				CreatedAt:    v.CreatedAt.Time,
			},
		)
	}

	return appUsers, nil
}
