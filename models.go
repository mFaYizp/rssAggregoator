package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/mFaYizp/rssagg/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreateAt  time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreateAt:  dbUser.CreateAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}
