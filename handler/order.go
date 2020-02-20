package handler

import (
	"hello/proto/order"
	"context"
	"fmt"
)

type Order struct {
}

func (e *Order) List(ctx context.Context, s *order.SayParam, r *order.SayResponse) error {
	fmt.Println(s.Msg)
	d := &order.Pair{
		Key:    23,
		Values: "xiaobnaijun",
	}
	var k = make(map[string]*order.Pair)
	k["xiaobai"] = d
	*r = order.SayResponse{
		Msg:    s.Msg,
		Values: []string{"xiaobai", "xiaoli"},
		Header: k,
		Type:   4,
	}
	fmt.Println("为什么我不能输出呢")
	return nil
}
