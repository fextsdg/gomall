package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall/app/checkout/conf"
	"gomall/common/clientsuite"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"
	"gomall/rpc_gen/kitex_gen/order/orderservice"
	"gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	Once               sync.Once
	ProductClient      productcatalogservice.Client
	CartClient         cartservice.Client
	PaymentClient      paymentservice.Client
	OrderClient        orderservice.Client
	CurrentServiceName = conf.GetConf().Kitex.Service
	RegistryAddress    = conf.GetConf().Registry.RegistryAddress[0]
)

func Init() {
	Once.Do(
		func() {
			initProductClient()
			initCartClient()
			initPaymentClient()
			initOrderClient()
		},
	)
}

func initProductClient() {
	suite := clientsuite.CommonClientSuite{
		CurrentServiceName: CurrentServiceName,
		RegistryAddress:    RegistryAddress,
	}
	ProductClient = productcatalogservice.MustNewClient("product", client.WithSuite(suite))
}

func initCartClient() {
	suite := clientsuite.CommonClientSuite{
		CurrentServiceName: CurrentServiceName,
		RegistryAddress:    RegistryAddress,
	}
	CartClient = cartservice.MustNewClient("cart", client.WithSuite(suite))
}
func initPaymentClient() {
	suite := clientsuite.CommonClientSuite{
		CurrentServiceName: CurrentServiceName,
		RegistryAddress:    RegistryAddress,
	}
	PaymentClient = paymentservice.MustNewClient("payment", client.WithSuite(suite))
}

func initOrderClient() {
	suite := clientsuite.CommonClientSuite{
		CurrentServiceName: CurrentServiceName,
		RegistryAddress:    RegistryAddress,
	}
	OrderClient = orderservice.MustNewClient("order", client.WithSuite(suite))
}

func InitTest1() {
	Once.Do(
		func() {
			initProductClientTest()
			initCartClientTest()
			initPaymentClientTest()
			initOrderClientTest()
		},
	)
}

func initProductClientTest() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatal(err)
	}
	ProductClient = productcatalogservice.MustNewClient("product", client.WithResolver(r))
}

func initCartClientTest() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatal(err)
	}
	CartClient = cartservice.MustNewClient("cart", client.WithResolver(r))
}
func initPaymentClientTest() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatal(err)
	}
	PaymentClient = paymentservice.MustNewClient("payment", client.WithResolver(r))
}
func initOrderClientTest() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		klog.Fatal(err)
	}
	OrderClient = orderservice.MustNewClient("order", client.WithResolver(r))
}
