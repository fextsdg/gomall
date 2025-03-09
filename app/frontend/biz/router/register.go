// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "gomall/app/frontend/biz/router/auth"
	cart "gomall/app/frontend/biz/router/cart"
	category "gomall/app/frontend/biz/router/category"
	checkout "gomall/app/frontend/biz/router/checkout"
	common "gomall/app/frontend/biz/router/common"
	home "gomall/app/frontend/biz/router/home"
	product "gomall/app/frontend/biz/router/product"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	checkout.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	auth.Register(r)

	common.Register(r)

	home.Register(r)
}
