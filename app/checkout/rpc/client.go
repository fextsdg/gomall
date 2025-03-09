package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall/app/checkout/conf"
	"gomall/rpc_gen/kitex_gen/cart/cartservice"
	"gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	Once          sync.Once
	ProductClient productcatalogservice.Client
	CartClient    cartservice.Client
	PaymentClient paymentservice.Client
)

func Init() {
	Once.Do(
		func() {
			initProductClient()
			initCartClient()
			initPaymentClient()
		},
	)
}

func initProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	ProductClient = productcatalogservice.MustNewClient("product", client.WithResolver(r))
}

func initCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	CartClient = cartservice.MustNewClient("cart", client.WithResolver(r))
}
func initPaymentClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		klog.Fatal(err)
	}
	PaymentClient = paymentservice.MustNewClient("payment", client.WithResolver(r))
}

func InitTest1() {
	Once.Do(
		func() {
			initProductClientTest()
			initCartClientTest()
			initPaymentClientTest()
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
