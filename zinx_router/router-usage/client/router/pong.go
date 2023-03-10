package router

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
)

// PongRouter MsgIdPong的路由
type PongRouter struct {
	znet.BaseRouter
}

//Ping Handle MsgIdPong的路由
func (r *PongRouter) Handle(request ziface.IRequest) {
	//读取客户端的数据
	fmt.Println("Handle: recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
}
