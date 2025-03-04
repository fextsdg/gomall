// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "gomall/app/frontend/biz/router/auth"
	common "gomall/app/frontend/biz/router/common"
	home "gomall/app/frontend/biz/router/home"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	auth.Register(r)

	common.Register(r)

	home.Register(r)
}
