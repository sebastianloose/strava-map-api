package cache

import (
	"errors"
	"time"

	"github.com/sebastianloose/strava-map-api/model"
)

var users map[int64]model.User

func AddUser(user model.User) {
	users[user.UserId] = user
}

func GetUser(id int64) (model.User, error) {
	user, exists := users[id]
	if exists {
		return user, nil
	}
	return model.User{}, errors.New("user not found")
}

func StartUserCacheWorker() {
	users = make(map[int64]model.User)

	for range time.Tick(time.Second * 1) {
		for userId, user := range users {
			if user.ExpiresAt < int(time.Now().Unix()) {
				delete(users, userId)
			}
		}
	}
}
