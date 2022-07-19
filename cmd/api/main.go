package main

import (
	"fmt"
	"net/http"
	"time"

	"MyDouyin/cmd/api/handlers"
	"MyDouyin/cmd/api/rpc"

	// jwt "github.com/appleboy/gin-jwt/v2"
	"MyDouyin/pkg/jwt"
	"MyDouyin/pkg/ttviper"
	"github.com/cloudwego/kitex/pkg/klog"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

var (
	Config      = ttviper.ConfigInit("TIKTOK_API", "apiConfig")
	ServiceName = Config.Viper.GetString("Server.Name")
	ServiceAddr = fmt.Sprintf("%s:%d", Config.Viper.GetString("Server.Address"), Config.Viper.GetInt("Server.Port"))
	EtcdAddress = fmt.Sprintf("%s:%d", Config.Viper.GetString("Etcd.Address"), Config.Viper.GetInt("Etcd.Port"))
	Jwt         *jwt.JWT
)

func Init() {
	rpc.InitRPC(&Config)
	Jwt = jwt.NewJWT([]byte(Config.Viper.GetString("JWT.signingKey")))
}

func main() {
	Init()
	r := gin.New()

	logger := Config.InitLogger()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(logger, time.RFC3339, false))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	douyin := r.Group("/douyin")
	user := douyin.Group("/user")
	user.POST("/login/", handlers.Login)
	user.POST("/register/", handlers.Register)
	user.GET("/", handlers.GetUserById)

	if err := http.ListenAndServe(ServiceAddr, r); err != nil {
		klog.Fatal(err)
	}
}
