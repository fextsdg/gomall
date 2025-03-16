package main

import (
	"gomall/app/email/biz/consumer"
	"gomall/app/email/infra/mq"
	"gomall/common/mtl"
	"gomall/common/servicesuite"
	"net"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"go.uber.org/zap/zapcore"
	"gomall/app/email/conf"
	"gomall/rpc_gen/kitex_gen/email/emailservice"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddress    = conf.GetConf().Registry.RegistryAddress[0]
)

func main() {
	mtl.InitMetrics(CurrentServiceName, conf.GetConf().Kitex.MetricsPort, RegistryAddress)
	mq.Init()       //初始化nats
	consumer.Init() //初始化消费者
	opts := kitexInit()

	svr := emailservice.NewServer(new(EmailServiceImpl), opts...)

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
