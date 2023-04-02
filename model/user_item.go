package model

import "github.com/google/uuid"

type User struct {
	UserId      uuid.UUID
	AccessToken string
	ExpiresAt   int
}
