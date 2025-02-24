package main

import (
	"github.com/joho/godotenv"
	"gomall/probuf/demo/biz/dal"
	"gomall/probuf/demo/biz/dal/mysql"
	"gomall/probuf/demo/biz/model"
	"gorm.io/gorm"
	"log"
)

func main() {
	//1.先加载环境变量，以供mysql.init使用
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln("预加载环境变量失败，err:=%v", err)
	}
	dal.Init()
	var DB *gorm.DB = mysql.DB
	//CURD
	//Create
	//err = DB.Create(&model.User{
	//	Email:    "demo@example.com",
	//	Password: "42316682712wssw",
	//}).Error
	//
	//if err != nil {
	//	log.Fatalf("创建数据失败，err:=%v", err)
	//}

	//update

	//err = DB.Model(&model.User{}).Where("email=?", "demo@example.com").Update("password", "1234567").Error
	//if err != nil {
	//	log.Fatalf("更新数据失败，err:=%v", err)
	//}

	//read
	var user model.User
	//res := DB.Model(&model.User{}).Where("email=?", "demo@example.com").Find(&user)
	//if res.Error != nil {
	//	log.Fatalf("获取数据失败，err:=%v", res.Error)
	//}
	//
	//log.Printf("获取%d条数据，数据：%v", res.RowsAffected, user)

	//delete

	res := DB.Where("email=?", "demo@example.com").Delete(&user)
	if res.Error != nil {
		log.Fatalf("删除数据失败，err:=%v", res.Error)
	}
	log.Printf("删除%d条数据", res.RowsAffected)

}
