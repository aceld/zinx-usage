package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type TestRouter struct {
	znet.BaseRouter
}

func (t *TestRouter) Handle(req ziface.IRequest) {
	fmt.Println("--> Call Handle")
}

func main() {
	s := znet.NewServer()
	s.AddRouter(1, &TestRouter{})
	s.Serve()
}
