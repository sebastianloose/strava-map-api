package cache

import (
	"github.com/sebastianloose/strava-map-api/model"
)

var activities map[int64]model.ActivityDetailedRoute

func AddDetailedActivity(activity model.ActivityDetailedRoute) {
	activities[activity.Id] = activity
}

func GetDetailedActivity(id int64, user model.User) (model.ActivityDetailedRoute, bool) {
	activity, exists := activities[id]
	if exists && activity.UserId == user.UserId {
		return activity, true
	}
	return model.ActivityDetailedRoute{}, false
}

func StartActivityCacheWorker() {
	activities = make(map[int64]model.ActivityDetailedRoute)
}
