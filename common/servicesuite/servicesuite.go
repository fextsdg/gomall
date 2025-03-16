package servicesuite

import (
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/monitor-prometheus"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall/common/mtl"
)

type CommonServiceSuite struct {
	CurrentServiceName string
	RegistryAddress    string
}

// 定义 GetOptions 方法，返回一个 server.Option 切片，用于配置服务器选项。
func (s CommonServiceSuite) Options() []server.Option {
	// 创建一个 server.Option 切片，用于存储服务器配置选项。
	opts := []server.Option{
		// 设置服务器的基本信息，包括服务名称。
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: s.CurrentServiceName}),
		// 使用 HTTP2 协议处理元数据。
		server.WithMetaHandler(transmeta.ServerHTTP2Handler),
		// 设置服务器的追踪器，使用 Prometheus 监控，并禁用服务器端的追踪。
		server.WithTracer(prometheus.NewServerTracer("", "", prometheus.WithDisableServer(true), prometheus.WithRegistry(mtl.Registry))),
	}

	//服务注册中心
	r, err := consul.NewConsulRegister(s.RegistryAddress)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithRegistry(r))
	// 返回配置好的服务器选项切片。
	return opts
}
