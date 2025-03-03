package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"

	user "gomall/rpc_gen/kitex_gen/user"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

var (
	err error
)

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.PasswordConfirm == "" {
		return nil, errors.New("用户名或密码不能为空！")
	}
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("输入的两次密码不一致！")
	}

	passwordHased, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newuser := model.User{Email: req.Email, PasswordHashed: string(passwordHased)}
	err = model.Create(mysql.DB, &newuser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResp{UserId: int32(newuser.ID)}, nil
}
