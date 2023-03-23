package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"time"
)

//客户端自定义业务
func pingLoop(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(1, []byte("Ping...Ping...Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			break
		}

		//获取链接携带的熟悉参数
		name, _ := conn.GetProperty("Name")
		home, _ := conn.GetProperty("Home")
		fmt.Println("Name = ", name, "Home = ", home)

		time.Sleep(1 * time.Second)
	}
}

//创建连接的时候执行
func onClientStart(conn ziface.IConnection) {
	fmt.Println("onClientStart is Called ... ")

	//给当前的IConnection设置携带属性参数
	conn.SetProperty("Name", "aceld")
	conn.SetProperty("Home", "https://www.yuque.com/aceld")

	go pingLoop(conn)
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
