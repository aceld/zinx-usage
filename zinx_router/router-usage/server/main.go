package main

import (
	"github.com/aceld/zinx/znet"
	"zinx-usage/zinx_router/router-usage/common"
	"zinx-usage/zinx_router/router-usage/server/router"
)

func main() {
	//1 创建一个server服务
	s := znet.NewServer()

	//2 配置路由
	s.AddRouter(common.MsgIdPing, &router.PingRouter{})

	//3 启动服务
	s.Serve()
}
