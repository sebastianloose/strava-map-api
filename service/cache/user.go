package cache

import (
	"errors"
	"time"

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

func StartCacheWorker() {
	for range time.Tick(time.Second * 1) {
		for i := 0; i < len(User); i++ {
			if User[i].ExpiresAt < int(time.Now().Unix()) {
				User = append(User[:i], User[i+1:]...)
				i--
			}
		}
	}
}

func AddUser(user model.User) {
	User = append(User, user)
}
