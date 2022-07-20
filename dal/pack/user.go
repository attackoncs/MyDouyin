/*
 * 封装 User 数据库数据为 RPC Server 端响应
 */

package pack

import (
	"context"
	"errors"

	"MyDouyin/kitex_gen/user"

	"gorm.io/gorm"

	"MyDouyin/dal/db"
)

// User pack user info
func User(ctx context.Context, u *db.User, fromID int64) (*user.User, error) {
	if u == nil {
		return &user.User{
			Name: "已注销用户",
		}, nil
	}

	follow_count := int64(u.FollowerCount)
	follower_count := int64(u.FollowerCount)

	isFollow := false
	relation, err := db.GetRelation(ctx, fromID, int64(u.ID))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if relation != nil {
		isFollow = true
	}
	return &user.User{
		Id:            int64(u.ID),
		Name:          u.UserName,
		FollowCount:   &follow_count,
		FollowerCount: &follower_count,
		IsFollow:      isFollow,
	}, nil
}

// Users pack list of user info
func Users(ctx context.Context, us []*db.User, fromID int64) ([]*user.User, error) {
	users := make([]*user.User, 0)
	for _, u := range us {
		user2, err := User(ctx, u, fromID)
		if err != nil {
			return nil, err
		}

		if user2 != nil {
			users = append(users, user2)
		}
	}
	return users, nil
}
