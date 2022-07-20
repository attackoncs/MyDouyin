/*
 * 初始化数据层
 */

package dal

import (
	db "MyDouyin/dal/db"
)

// Init init dal
func Init() {
	db.Init() // mysql init
}
