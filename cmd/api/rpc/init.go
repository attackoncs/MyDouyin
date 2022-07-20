/*
 * 基于配置信息初始化 RPC 客户端
 */

package rpc

import "MyDouyin/pkg/ttviper"

// InitRPC init rpc client
func InitRPC(Config *ttviper.Config) {
	UserConfig := ttviper.ConfigInit("TIKTOK_USER", "userConfig")
	initUserRpc(&UserConfig)

	FeedConfig := ttviper.ConfigInit("TIKTOK_FEED", "feedConfig")
	initFeedRpc(&FeedConfig)

	PublishConfig := ttviper.ConfigInit("TIKTOK_PUBLISH", "publishConfig")
	initPublishRpc(&PublishConfig)

	FavoriteConfig := ttviper.ConfigInit("TIKTOK_FAVORITE", "favoriteConfig")
	initFavoriteRpc(&FavoriteConfig)

	CommentConfig := ttviper.ConfigInit("TIKTOK_COMMENT", "commentConfig")
	initCommentRpc(&CommentConfig)

	RelationConfig := ttviper.ConfigInit("TIKTOK_RELATION", "relationConfig")
	initRelationRpc(&RelationConfig)
}
