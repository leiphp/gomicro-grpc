package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/consul"
	"gomicro-grpc/ServiceImpl"
	"gomicro-grpc/Services"
)

func main(){
	consulReg := consul.NewRegistry(
		registry.Addrs("192.168.1.104:8500"),
	)

	prodService := micro.NewService(
		micro.Name("prodservice"),
		micro.Address(":8011"),
		micro.Registry(consulReg),
		)
	prodService.Init()
	Services.RegisterProdServiceHandler(prodService.Server(),new(ServiceImpl.ProdService))
	prodService.Run()
}
