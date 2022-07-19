package pack

import (
	"MyDouyin/kitex_gen/user"

	"MyDouyin/dal/db"
)

// User pack user info
func User(u *db.User) *user.User {
	if u == nil {
		return nil
	}

	follow_count := int64(u.FollowerCount)
	follower_count := int64(u.FollowerCount)

	return &user.User{Id: int64(u.ID), Name: u.UserName, FollowCount: &follow_count, FollowerCount: &follower_count}
}

// Users pack list of user info
func Users(us []*db.User) []*user.User {
	users := make([]*user.User, 0)
	for _, u := range us {
		if user2 := User(u); user2 != nil {
			users = append(users, user2)
		}
	}
	return users
}
