package mtl

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/server"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net"
	"net/http"
)

var (
	Registry *prometheus.Registry
)

// 定义 InitMetrics 函数，用于初始化 Prometheus 监控和 Consul 注册功能。
// 参数说明：
// - serviceName: 服务名称，用于标识当前服务。
// - metricsPort: Prometheus 指标暴露的端口。
// - registryAddr: Consul 注册地址。
func InitMetrics(serviceName, metricsPort, registryAddr string) (registry.Registry, *registry.Info) {
	// 创建一个新的 Prometheus 注册表实例，用于管理所有指标。
	Registry = prometheus.NewRegistry()

	// 注册 Go 运行时相关的指标收集器，例如 goroutine 数量、堆内存等。
	Registry.MustRegister(collectors.NewGoCollector())

	// 注册进程相关的指标收集器，例如 CPU 使用率、内存使用情况等。
	Registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))

	// 初始化 Consul 注册器，用于将服务注册到 Consul 中。
	r, _ := consul.NewConsulRegister(registryAddr)

	// 将 metricsPort 转换为 TCP 地址格式，以便后续使用。
	addr, _ := net.ResolveTCPAddr("tcp", metricsPort)

	// 定义服务注册信息，包括服务名称、地址、权重和标签。
	// 注意：此处的服务名称固定为 "prometheus"，实际应根据需求调整。
	registryInfo := &registry.Info{
		ServiceName: "prometheus",
		Addr:        addr,
		Weight:      1,
		Tags:        map[string]string{"service": serviceName},
	}

	// 将服务信息注册到 Consul 中。
	r.Register(registryInfo)

	// 注册一个关闭钩子，在服务停止时自动注销服务。
	server.RegisterShutdownHook(
		func() {
			r.Deregister(registryInfo)
		})

	// 配置 HTTP 路由，将 "/metrics" 路径映射到 Prometheus 的指标处理器。
	http.Handle("/metrics", promhttp.HandlerFor(Registry, promhttp.HandlerOpts{}))

	// 启动一个 Goroutine，监听指定的 metricsPort 端口，提供 Prometheus 指标服务。
	go http.ListenAndServe(metricsPort, nil)

	return r, registryInfo
}
