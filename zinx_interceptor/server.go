package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

type MyInterceptor struct{}

func (m *MyInterceptor) Intercept(chain ziface.IChain) ziface.IcResp {
	request := chain.Request()
	// 这一层是自定义拦截器处理逻辑，这里只是简单打印输入
	iRequest := request.(ziface.IRequest)
	//zlog.Ins().InfoF("自定义拦截器, 收到消息：%s", iRequest.GetData())
	fmt.Println("自定义拦截器, 收到消息：%s", iRequest.GetData())
	return chain.Proceed(chain.Request())
}

type HelloRouter struct {
	znet.BaseRouter
}

func (hr *HelloRouter) Handle(request ziface.IRequest) {
	//zlog.Ins().InfoF(string(request.GetData()))
	fmt.Println(string(request.GetData()))
}

func main() {
	// 创建server 对象
	server := znet.NewServer()
	// 添加路由映射
	server.AddRouter(1, &HelloRouter{})
	// 添加自定义拦截器
	server.AddInterceptor(&MyInterceptor{})
	// 启动服务
	server.Serve()
}
