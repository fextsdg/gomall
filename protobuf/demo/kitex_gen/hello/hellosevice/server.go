// Code generated by Kitex v0.9.1. DO NOT EDIT.
package hellosevice

import (
	server "github.com/cloudwego/kitex/server"
	hello "gomall/probuf/demo/kitex_gen/hello"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler hello.HelloSevice, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}

func RegisterService(svr server.Server, handler hello.HelloSevice, opts ...server.RegisterOption) error {
	return svr.RegisterService(serviceInfo(), handler, opts...)
}
