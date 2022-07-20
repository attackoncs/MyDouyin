/*
 * 获取发布视频列表 操作业务逻辑
 */

package command

import (
	"context"

	"MyDouyin/dal/pack"
	"MyDouyin/kitex_gen/feed"
	"MyDouyin/kitex_gen/publish"

	"MyDouyin/dal/db"
)

type PublishListService struct {
	ctx context.Context
}

// NewPublishListService new PublishListService
func NewPublishListService(ctx context.Context) *PublishListService {
	return &PublishListService{ctx: ctx}
}

// PublishList publish video.
func (s *PublishListService) PublishList(req *publish.DouyinPublishListRequest) (vs []*feed.Video, err error) {
	videos, err := db.PublishList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	vs, err = pack.Videos(s.ctx, videos, &req.UserId)
	if err != nil {
		return nil, err
	}

	return vs, nil
}
