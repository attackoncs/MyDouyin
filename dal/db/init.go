/*
 * 数据库初始化
 */

package db

import (
	"fmt"
	"time"

	"MyDouyin/pkg/dlog"
	"MyDouyin/pkg/ttviper"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var (
	DB     *gorm.DB
	Config = ttviper.ConfigInit("TIKTOK_DB", "dbConfig")
)

// InitDB Init init DB
func InitDB() {
	var err error

	logger := zapgorm2.New(dlog.InitLog())
	logger.SetAsDefault() // optional: configure gorm to use this zapgorm.Logger for callbacks

	viper := Config.Viper
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=%s",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
		viper.GetString("mysql.charset"),
		viper.GetBool("mysql.parseTime"),
		viper.GetString("mysql.loc"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			// PrepareStmt: true,
			// SkipDefaultTransaction: true,
			Logger: logger,
		},
	)
	if err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	if err = DB.Use(otelgorm.NewPlugin()); err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	if err := DB.AutoMigrate(&User{}, &Video{}, &Comment{}, &Relation{}); err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		logger.ZapLogger.Panic(err.Error())
	}

	if err := sqlDB.Ping(); err != nil {
		logger.ZapLogger.Panic(err.Error())
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func Init() {
	InitDB()
}
