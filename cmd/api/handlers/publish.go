/*
 * @Description: 定义 Publish API 的 handler
 */

package handlers

import (
	"bytes"
	"context"
	"io"
	"strconv"

	"MyDouyin/pkg/errno"

	"MyDouyin/dal/pack"
	"MyDouyin/kitex_gen/publish"

	"MyDouyin/cmd/api/rpc"

	"github.com/gin-gonic/gin"
)

// PublishAction 传递 发布视频操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishAction(c *gin.Context) {
	var paramVar PublishActionParam
	token := c.PostForm("token")
	title := c.PostForm("title")

	file, _, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ErrDecodingFailed))
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		SendResponse(c, pack.BuildPublishResp(err))
		return
	}

	paramVar.Token = token
	paramVar.Title = title

	resp, err := rpc.PublishAction(context.Background(), &publish.DouyinPublishActionRequest{
		Title: paramVar.Title,
		Token: paramVar.Token,
		Data:  buf.Bytes(),
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

// PublishList 传递 获取视频列表操作 的上下文至 Publish 服务的 RPC 客户端, 并获取相应的响应.
func PublishList(c *gin.Context) {
	var paramVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}
	paramVar.UserId = int64(userid)
	paramVar.Token = c.Query("token")

	if len(paramVar.Token) == 0 || paramVar.UserId < 0 {
		SendResponse(c, pack.BuildPublishListResp(errno.ErrBind))
		return
	}

	resp, err := rpc.PublishList(context.Background(), &publish.DouyinPublishListRequest{
		UserId: paramVar.UserId,
		Token:  paramVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuildPublishListResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
