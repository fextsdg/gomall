package mysql

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gomall/app/order/biz/model"
	"gomall/app/order/conf"
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
	if !DB.Migrator().HasTable("order") {
		klog.Info("Order表不存在，正在迁移...")
		if err = DB.AutoMigrate(&model.Order{}); err != nil {
			panic(err)
		}

	}
	if !DB.Migrator().HasTable("order_item") {
		klog.Info("order_item表不存在，正在迁移...")
		if err = DB.AutoMigrate(&model.OrderItem{}); err != nil {
			panic(err)
		}

	}
}

func InitTest1() {
	dsn := "%s:%s@tcp(%s:3306)/order?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(fmt.Sprintf(dsn, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if !DB.Migrator().HasTable("order") {
		klog.Info("Order表不存在，正在迁移...")
		if err = DB.AutoMigrate(&model.Order{}); err != nil {
			panic(err)
		}

	}
	if !DB.Migrator().HasTable("order_item") {
		klog.Info("order_item表不存在，正在迁移...")
		if err = DB.AutoMigrate(&model.OrderItem{}); err != nil {
			panic(err)
		}

	}
}
