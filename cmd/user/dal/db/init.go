package db

import (
	"fmt"

	"MyDouyin/pkg/ttviper"
	"github.com/spf13/pflag"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	TableName = pflag.String("mysql.tablename", "", "Mysql Table Name")
)

// Init init DB
func InitDB(config *ttviper.Config) {
	var err error
	viper := config.Viper
	*TableName = viper.GetString("mysql.tablename")
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
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}

	if err := DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}

func Init(config *ttviper.Config) {
	InitDB(config)
}
