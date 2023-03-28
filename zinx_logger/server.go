package main

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
)

type TestRouter struct {
	znet.BaseRouter
}

// Handle -
func (t *TestRouter) Handle(req ziface.IRequest) {
	zlog.Ins().InfoF("Call Handle")
	if err := req.GetConnection().SendMsg(0, []byte("test2")); err != nil {
		zlog.Ins().ErrorF("err: %v", err)
	}
}

func main() {
	s := znet.NewServer()

	s.AddRouter(1, &TestRouter{})

	//设置日志
	zlog.SetLogger(new(MyLogger))

	s.Serve()
}
