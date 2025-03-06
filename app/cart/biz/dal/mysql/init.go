package mysql

import (
	"fmt"
	"gomall/app/cart/conf"
	"gomall/app/cart/model"
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
	if !DB.Migrator().HasTable("cart") {
		fmt.Println("购物车表不存在，正在创建...")
		err = DB.AutoMigrate(&model.Cart{})
		if err != nil {
			panic(err)
		}
	}

}

func InitTest() {

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:3306)/cart?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"))),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if !DB.Migrator().HasTable("cart") {
		fmt.Println("购物车表不存在，正在创建...")
		err = DB.AutoMigrate(&model.Cart{})
		if err != nil {
			panic(err)
		}
	}

}
