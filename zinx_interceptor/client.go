package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"time"
)

/*
func main() {
	// 客户端向服务端发送消息
	client := znet.NewClient("127.0.0.1", 8999)

	client.SetOnConnStart(func(connection ziface.IConnection) {
		err := connection.SendMsg(1, []byte("hello zinx"))
		if err != nil {
			fmt.Println(err)
		}
	})

	client.Start()

	select {}
}

*/

//客户端自定义业务
func pingLoop(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(1, []byte("Ping...Ping...Ping...[FromClient]"))
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
	/*
		err := conn.SendMsg(1, []byte("Ping...Ping...Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
		}
	*/
}

func main() {
	//创建Client客户端
	client := znet.NewClient("127.0.0.1", 8999)

	//设置链接建立成功后的钩子函数
	client.SetOnConnStart(onClientStart)

	//启动客户端
	client.Start()

	//防止进程退出，等待中断信号
	select {}
}
