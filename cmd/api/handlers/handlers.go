package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SendResponse pack response
func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

type UserRegisterParam struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserParam struct {
	UserId int64  `json:"user_id,omitempty"` // 用户id
	Token  string `json:"token,omitempty"`   // 用户鉴权token
}
