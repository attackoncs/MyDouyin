/*
 * 封装 FavoriteVideos 数据库数据为 RPC Server 端响应
 */

package pack

import (
	"context"

	"MyDouyin/kitex_gen/feed"

	"MyDouyin/dal/db"
)

// FavoriteVideos pack favoriteVideos info.
func FavoriteVideos(ctx context.Context, vs []db.Video, uid *int64) ([]*feed.Video, error) {
	videos := make([]*db.Video, 0)
	for _, v := range vs {
		videos = append(videos, &v)
	}

	packVideos, err := Videos(ctx, videos, uid)
	if err != nil {
		return nil, err
	}

	return packVideos, nil
}
