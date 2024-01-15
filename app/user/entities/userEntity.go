package entities

import "time"

type AppUser struct {
	Id           uint32
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}
