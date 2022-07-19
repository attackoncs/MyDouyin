package dal

import (
	db "MyDouyin/cmd/user/dal/db"
	"MyDouyin/pkg/ttviper"
)

// Init init dal
func Init(config *ttviper.Config) {
	db.Init(config) // mysql init
}
