package main

import (
	"github.com/joho/godotenv"
	"gomall/app/cart/biz/dal"
	"gomall/app/cart/rpc"
	"gomall/common/mtl"
	"gomall/common/servicesuite"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gomall/app/cart/conf"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddress    = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	err1 := godotenv.Load(".env")
	if err1 != nil {
		panic(err1)
	}
	//初始化prometheus检测指标
	mtl.InitMetrics(CurrentServiceName, conf.GetConf().Kitex.MetricsAddr, RegistryAddress)
	opts := kitexInit()
	dal.Init()
	rpc.Init() //初始化商品服务
	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))
	suite := servicesuite.CommonServiceSuite{
		CurrentServiceName: CurrentServiceName,
		RegistryAddress:    RegistryAddress,
	}
	opts = append(opts, server.WithSuite(suite))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
