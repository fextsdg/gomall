package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gomall/app/user/biz/dal/mysql"
	"gomall/app/user/biz/model"
	user "gomall/rpc_gen/kitex_gen/user"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginReq) (resp *user.LoginResp, err error) {
	// Finish your business logic.
	if req.Email == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空！")
	}

	u, err := model.GetUserByEmail(mysql.DB, req.GetEmail())
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误！")
	}

	return &user.LoginResp{UserId: int32(u.ID)}, err
}
