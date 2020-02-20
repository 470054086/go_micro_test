package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"hello/handler"
	"hello/subscriber"
	example "hello/proto/example"
	"github.com/micro/go-plugins/registry/etcd"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.srv.hello"),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry()),
	)

	// Initialise service
	service.Init(
	)

	// Register Handler
	example.RegisterExampleHandler(service.Server(), new(handler.Example))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("com.foo.srv.hello", service.Server(), new(subscriber.Example))

	// Register Function as Subscriber
	micro.RegisterSubscriber("com.foo.srv.hello", service.Server(), subscriber.Handler)



	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
