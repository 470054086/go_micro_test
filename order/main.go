package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-log"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"hello/proto/order"
	"hello/handler"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.srv.order"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(func(op *registry.Options) {
			op.Addrs = []string{
				"http://192.168.38.129:2379",
			}
		})),
	)

	// Initialise service
	service.Init()
	// Register Handler
	order.RegisterOrderHandler(service.Server(), new(handler.Order))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
