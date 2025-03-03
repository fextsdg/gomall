package service

import (
	"context"
	"github.com/joho/godotenv"
	"gomall/app/user/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	err = godotenv.Load("../../.env")
	if err != nil {

		panic(err)
	}
	mysql.Init()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{
		Email:    "234221323@qeerw.com",
		Password: "1234567", //123456
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
