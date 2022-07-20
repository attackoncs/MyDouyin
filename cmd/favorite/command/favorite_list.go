package command

import (
	"context"

	"MyDouyin/dal/pack"
	"MyDouyin/kitex_gen/favorite"
	"MyDouyin/kitex_gen/feed"

	"MyDouyin/dal/db"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context) *FavoriteListService {
	return &FavoriteListService{
		ctx: ctx,
	}
}

func (s *FavoriteListService) FavoriteList(req *favorite.DouyinFavoriteListRequest) ([]*feed.Video, error) {
	FavoriteVideos, err := db.FavoriteList(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	videos, err := pack.FavoriteVideos(s.ctx, FavoriteVideos, &req.UserId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}
