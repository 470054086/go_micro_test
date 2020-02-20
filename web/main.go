package main

import (
	"github.com/micro/go-micro/web"
	"github.com/micro/go-micro/registry/etcd"
	"hello/proto/order"
	"fmt"
	"context"
	"net/http"
	"github.com/micro/go-micro"
)
var srcClient order.OrderService
//定义一个web服务
func main()  {
	service := web.NewService(
		web.Name("com.foo.hello.web"),
		web.Registry(etcd.NewRegistry()),
	)
	service.Init()

	serviceClient := micro.NewService(
		micro.Name("com.foo.hello.client"),
		micro.Registry(etcd.NewRegistry()),
	)
	serviceClient.Init()
	//构建客户端
	srcClient = order.NewOrderService("com.foo.srv.order",serviceClient.Client())

	//web服务器启动
	service.HandleFunc("/abc", orderService)

	service.Run()

}

func orderService(writer http.ResponseWriter, request *http.Request) {
	msg := request.URL.Query().Get("msg")

	params := order.SayParam{
		Msg:msg,
	}
	r,err := srcClient.List(context.Background(),&params)
	fmt.Fprint(writer,r)
	fmt.Fprint(writer,err)

}