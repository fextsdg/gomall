// Code generated by hertz generator.

package main

import (
	"context"
	"fmt"
	"github.com/hertz-contrib/sessions"
	"github.com/hertz-contrib/sessions/redis"
	"github.com/joho/godotenv"
	"gomall/app/frontend/middleware"
	"os"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"go.uber.org/zap/zapcore"
	"gomall/app/frontend/biz/router"
	"gomall/app/frontend/conf"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	//加载环境变量
	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("加载环境变量出错,err:=%v", err))
	}
	// init dal
	// dal.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))

	registerMiddleware(h)
	middleware.Register(h)
	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.GET("/sign-in", func(c context.Context, ctx *app.RequestContext) {
		uh := utils.H{
			"Name": "Sign-In",
			"Next": ctx.Request.Header.Get("Referer"), //或取上一个页面来源
		}

		ctx.HTML(consts.StatusOK, "sign-in", uh) //渲染静态页面
	})
	//注册路由
	h.GET("/sign-up", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "sign-up", utils.H{"Name": "Sign-Up"})
	})

	router.GeneratedRegister(h)
	h.LoadHTMLGlob("template/*") //指定html模板文件
	h.Static("/static", "./")    //查找静态资源
	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	hlog.SetOutput(asyncWriter)
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})

	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	//session
	store, _ := redis.NewStore(10, "tcp", conf.GetConf().Redis.Address, "", []byte(os.Getenv("SESSION_SECRET")))
	h.Use(sessions.New("fextsdg-shop", store)) //fextsdg-shop 是会话的名称，它用于标识存储在 Redis 中的会话数据。实际存储的是与该会话相关的数据。
}
