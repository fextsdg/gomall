package clientsuite

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

type CommonClientSuite struct {
	CurrentServiceName string
	RegistryAddress    string
}

// 定义 GetOptions 方法，返回一个 client.Option 切片，用于配置客户端选项。
func (s CommonClientSuite) Options() []client.Option {
	// 创建一个 client.Option 切片，用于存储客户端配置选项。
	opts := []client.Option{
		// 使用 HTTP2 协议处理元数据。
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		// 设置客户端的基本信息，包括服务名称。
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: s.CurrentServiceName}),
		// 设置传输协议为 gRPC。
		client.WithTransportProtocol(transport.GRPC),
	}
	r, err := consul.NewConsulResolver(s.RegistryAddress)

	if err != nil {
		panic(err)
	}
	opts = append(opts, client.WithResolver(r))
	// 返回配置好的客户端选项切片。
	return opts
}
