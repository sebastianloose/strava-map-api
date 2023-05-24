package cache

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/sebastianloose/strava-map-api/model"
)

var users map[uuid.UUID]model.User

func AddUser(user model.User) {
	users[user.UserId] = user
}

func GetUser(id uuid.UUID) (model.User, error) {
	user, exists := users[id]
	if exists {
		return user, nil
	}
	return model.User{}, errors.New("user not found")
}

func StartCacheWorker() {
	users = make(map[uuid.UUID]model.User)

	for range time.Tick(time.Second * 1) {
		for userId, user := range users {
			if user.ExpiresAt < int(time.Now().Unix()) {
				delete(users, userId)
			}
		}
	}
}
