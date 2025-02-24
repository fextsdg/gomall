package dal

import (
	"gomall/probuf/demo/biz/dal/mysql"
	"gomall/probuf/demo/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
