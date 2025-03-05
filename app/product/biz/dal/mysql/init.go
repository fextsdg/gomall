package mysql

import (
	"fmt"
	"github.com/cloudwego/kitex/tool/internal_pkg/log"
	"gomall/app/product/conf"
	"gomall/app/product/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := conf.GetConf().MySQL.DSN

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(dsn, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	// 检查 Product 表是否存在
	if !DB.Migrator().HasTable(&model.Product{}) {
		// 如果表不存在，则创建表
		log.Info("Product 表不存在,正在迁移...")
		err = DB.AutoMigrate(&model.Product{})
		if err != nil {
			panic(err)
		}
	}

	// 检查 Category 表是否存在
	if !DB.Migrator().HasTable(&model.Category{}) {
		// 如果表不存在，则创建表
		log.Info("Category 表不存在,正在迁移...")
		err = DB.AutoMigrate(&model.Category{})
		if err != nil {
			panic(err)
		}
	}

}
