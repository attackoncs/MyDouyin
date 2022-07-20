//定义 User API 的 handler
package handlers

import (
	"MyDouyin/dal/pack"
	"context"
	"strconv"

	"MyDouyin/pkg/errno"

	"MyDouyin/cmd/api/rpc"
	"MyDouyin/kitex_gen/user"

	"github.com/gin-gonic/gin"
)

// Register register user info
func Register(c *gin.Context) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Register(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func Login(c *gin.Context) {
	var registerVar UserRegisterParam
	registerVar.UserName = c.Query("username")
	registerVar.PassWord = c.Query("password")

	if len(registerVar.UserName) == 0 || len(registerVar.PassWord) == 0 {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ErrBind))
		return
	}

	resp, err := rpc.Login(context.Background(), &user.DouyinUserRegisterRequest{
		Username: registerVar.UserName,
		Password: registerVar.PassWord,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserRegisterResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}

func GetUserById(c *gin.Context) {
	var userVar UserParam
	userid, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		SendResponse(c, pack.BuilduserUserResp(errno.ErrBind))
		return
	}
	userVar.UserId = int64(userid)
	userVar.Token = c.Query("token")

	if len(userVar.Token) == 0 || userVar.UserId < 0 {
		SendResponse(c, pack.BuilduserUserResp(errno.ErrBind))
		return
	}

	resp, err := rpc.GetUserById(context.Background(), &user.DouyinUserRequest{
		UserId: userVar.UserId,
		Token:  userVar.Token,
	})
	if err != nil {
		SendResponse(c, pack.BuilduserUserResp(errno.ConvertErr(err)))
		return
	}
	SendResponse(c, resp)
}
