package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

type SessionUserIdKey string

const SessionUserId SessionUserIdKey = "user_id"

// 为了方便业务逻辑获取用户信息，从session中获取用户信息，然后放在context中
// GlobalAuth 返回一个中间件处理函数，用于全局认证。
// 该中间件主要负责将当前请求的用户ID从会话中取出，并将其添加到请求上下文中。
// 这样做是为了让后续的处理函数能够方便地访问用户ID，而无需再次从会话中获取。
func GlobalAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 获取默认会话，会话管理由gorilla/sessions库提供支持
		session := sessions.Default(ctx)
		// 将用户ID从会话中取出，并添加到请求上下文中
		// 这里的SessionUserId是一个全局变量或者常量，用于在上下文中标识用户ID的键
		c = context.WithValue(c, SessionUserId, session.Get("user_id"))
		// 调用下一个处理函数，将处理权移交给链中的下一个中间件或最终的处理函数
		ctx.Next(c)
	}
}

// Auth 是一个中间件函数，用于验证用户是否已登录
// 该函数符合 app.HandlerFunc 类型，可以作为路由处理函数
func Auth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 从请求上下文中获取用户ID，用于验证用户登录状态
		userid := c.Value(SessionUserId)
		// 用户未登录，则跳转登录界面
		if userid == nil {
			// 使用 HTTP 302 临时重定向到登录页面，并将当前请求路径作为登录后的跳转目标
			ctx.Redirect(consts.StatusFound, []byte("/sign-in?next="+ctx.FullPath()))
			// 中断当前请求处理
			ctx.Abort()
			return
		}
		// 用户已登录，继续执行下一个中间件或最终的处理函数
		ctx.Next(c)
	}
}
