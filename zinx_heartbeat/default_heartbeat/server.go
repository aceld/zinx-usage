package main

import (
	"github.com/aceld/zinx/znet"
	"time"
)

func main() {
	//创建Server
	s := znet.NewServer()

	//启动心跳检测，默认3秒检测一次
	s.StartHeartBeat(3 * time.Second)

	//启动服务
	s.Serve()
}
