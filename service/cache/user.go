package cache

import (
	"errors"

	"github.com/google/uuid"
	"github.com/sebastianloose/strava-map-api/model"
)

var User []model.User

func GetUserById(id uuid.UUID) (model.User, error) {
	for _, u := range User {
		if u.UserId == id {
			return u, nil
		}
	}
	return model.User{}, errors.New("user not found")
}

func AddUser(user model.User) {
	User = append(User, user)
}
