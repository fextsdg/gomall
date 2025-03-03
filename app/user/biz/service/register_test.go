package service

import (
	"context"
	"github.com/joho/godotenv"
	"gomall/app/user/biz/dal/mysql"
	user "gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	ctx := context.Background()
	err := godotenv.Load("../../.env")

	mysql.Init()
	if err != nil {
		panic(err)
	}
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{
		Email:           "",
		Password:        "1234567",
		PasswordConfirm: "123456",
	}
	resp, err := s.Run(req)
	if err != nil {
		t.Logf("err: %v", err)
	}

	if resp != nil {
		t.Logf("resp: %v", resp)
	}

	// todo: edit your unit test

}
