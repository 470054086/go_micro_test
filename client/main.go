package main

import (
	"github.com/micro/go-micro"
	example "hello/proto/example"
	"context"
	"fmt"
	"github.com/micro/go-micro/registry/etcd"
	"hello/proto/order"
)

//客户端代码
func main()  {
	service := micro.NewService(
		micro.Name("com.foo.hello.client"),
		micro.Registry(etcd.NewRegistry()),
	)
	//初始化
	service.Init()
	//生成客户端
	client := example.NewExampleService("com.foo.srv.hello",service.Client())
	d := example.Request{
		Name:"xiaobai",
	}
	res,err := client.Call(context.Background(),&d)
	fmt.Println(res,err)
	service1 := micro.NewService(
		micro.Name("com.foo.order.client"),
		micro.Registry(etcd.NewRegistry()),
	)

	orderClient := order.NewOrderService("com.foo.srv.order",service1.Client())
	params := order.SayParam{
		Msg:"hello,world",
	}
	r,err := orderClient.List(context.Background(),&params)
	fmt.Println(r,err)
}
