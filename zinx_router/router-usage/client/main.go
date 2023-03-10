package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"time"
	"zinx-usage/zinx_router/router-usage/client/router"
	"zinx-usage/zinx_router/router-usage/common"
)

//客户端自定义业务
func pingLoop(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(common.MsgIdPing, []byte("Ping...Ping...Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

//创建连接的时候执行
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")
	go pingLoop(conn)
}

func main() {
	//创建Client客户端
	client := znet.NewClient("127.0.0.1", 8999)

	//设置链接建立成功后的钩子函数
	client.SetOnConnStart(onClientStart)

	//设置消息读取路由
	client.AddRouter(common.MsgIdPong, &router.PongRouter{})

	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}
