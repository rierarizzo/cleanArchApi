package repositories

import (
	appError "cleanArchApi/app/error"
	"cleanArchApi/app/user/entities"
	"cleanArchApi/database/postgres/sqlc"
	"context"
	"database/sql"
	"errors"
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
	userRows, err := queries.GetUsers(context.TODO())

	if err != nil {
		slog.Error(fmt.Sprintf("SelectAppUsersData: %v", err))

		if errors.Is(err, sql.ErrNoRows) {
			return appUsers, nil
		}

		return appUsers, appError.ErrUnknown
	}

	for _, row := range userRows {
		appUsers = append(
			appUsers, entities.AppUser{
				Id:           uint32(row.AppUserID),
				Username:     row.Username,
				Email:        row.Email,
				PasswordHash: row.PasswordHash,
				CreatedAt:    row.CreatedAt.Time,
			},
		)
	}

	return appUsers, nil
}
