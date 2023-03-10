package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type TestRouter struct {
	znet.BaseRouter
}

func (t *TestRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("--> Call PreHandle")

	//跳转
	//req.Goto(znet.POST_HANDLE)
}

func (t *TestRouter) Handle(req ziface.IRequest) {
	fmt.Println("--> Call Handle")
	//终止
	req.Abort()
}

func (t *TestRouter) PostHandle(req ziface.IRequest) {
	fmt.Println("--> Call PostHandle")
}

func main() {
	s := znet.NewServer()
	s.AddRouter(1, &TestRouter{})
	s.Serve()
}
