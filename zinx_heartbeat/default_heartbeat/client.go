package main

import (
	"github.com/aceld/zinx/znet"
	"time"
)

func main() {
	//创建Client客户端
	client := znet.NewClient("127.0.0.1", 8999)

	//启动心跳检测、默认3秒检测一次
	client.StartHeartBeat(3 * time.Second)

	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}
