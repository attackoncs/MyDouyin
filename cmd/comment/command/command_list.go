/*
 * 获取评论列表 操作业务逻辑
 */

package command

import (
	"context"

	"MyDouyin/dal/pack"
	"MyDouyin/kitex_gen/comment"

	"MyDouyin/dal/db"
)

type CommentListService struct {
	ctx context.Context
}

// NewCommentListService NewCommentActionService new CommentActionService
func NewCommentListService(ctx context.Context) *CommentListService {
	return &CommentListService{
		ctx: ctx,
	}
}

// CommentList return comment list
func (s *CommentListService) CommentList(req *comment.DouyinCommentListRequest, fromID int64) ([]*comment.Comment, error) {
	Comments, err := db.GetVideoComments(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}

	comments, err := pack.Comments(s.ctx, Comments, fromID)
	if err != nil {
		return nil, err
	}
	return comments, nil
}
