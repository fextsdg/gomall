package rpc

import (
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall/app/frontend/conf"
	"gomall/app/frontend/utils"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"
	"gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"gomall/rpc_gen/kitex_gen/user/userservice"
	"sync"
)

var (
	once          sync.Once //保证只被初始化一次
	UserClient    userservice.Client
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
)

// 初始化客户端
func Init() {
	once.Do(func() {
		initUserClient()
		initProductClient()
		initCartClient()
	})
}

// initUserClient 初始化用户服务客户端。
//
// 该函数通过配置中的注册地址创建一个新的 Consul 解析器，
// 并使用该解析器创建用户服务客户端。如果在创建解析器时发生错误，
// 则记录错误并终止程序。
//
// 重要代码块说明：
// - 创建 Consul 解析器：根据配置文件中的注册地址初始化 Consul 解析器。
// - 错误处理：如果解析器创建失败，记录错误并退出程序。
// - 创建用户服务客户端：使用 Consul 解析器初始化用户服务客户端。
func initUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress[0])
	utils.MustHandlerError(err)
	UserClient = userservice.MustNewClient("user", client.WithResolver(r))
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress[0])
	utils.MustHandlerError(err)
	ProductClient = productcatalogservice.MustNewClient("product", client.WithResolver(r))
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Hertz.RegistryAddress[0])
	utils.MustHandlerError(err)
	CartClient = cartservice.MustNewClient("cart", client.WithResolver(r))
}
