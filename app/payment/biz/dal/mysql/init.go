package mysql

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"gomall/app/payment/conf"
	"gomall/app/payment/model"
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
	if !DB.Migrator().HasTable("payment_log") {
		klog.Info("payment_log表不存在，正在迁移...")
		err = DB.AutoMigrate(&model.PaymentLog{})
		if err != nil {
			panic(err)
		}
	}
}

func InitTest() {

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/payment?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if !DB.Migrator().HasTable("payment_log") {
		klog.Info("payment_log表不存在，正在迁移...")
		err = DB.AutoMigrate(&model.PaymentLog{})
		if err != nil {
			panic(err)
		}
	}

}
