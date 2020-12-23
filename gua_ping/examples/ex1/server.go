package main

import (
	"fmt"
	"go_basic/gua_ping/itf"
	"go_basic/gua_ping/net"
)

//自定义路由

type PingRouter struct {
	net.BaseRouter
}

func (router *PingRouter) Handle(request itf.IRequest) {
	fmt.Println("Call PingRouter Handle")

	err := request.GetConnection().SendBuffMsg(0, []byte("test...ping"))
	if err != nil {
		fmt.Println("ping router handle err:", err)
	}
}

//启动server端
func main() {
	server := net.NewServer()

	//添加路由
	server.AddRouter(0, &PingRouter{})
	server.Serve()
}
