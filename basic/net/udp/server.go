package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建 服务器 UDP 地址结构。指定 IP + port
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr err:", err)
		return
	}
	// 监听 客户端连接
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		fmt.Println("net.ListenUDP err:", err)
		return
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, raddr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("conn.ReadFromUDP err:", err)
			return
		}
		fmt.Printf("接收到客户端[%s]：%s", raddr, string(buf[:n]))

		conn.WriteToUDP([]byte("I-AM-SERVER"), raddr) // 简单回写数据给客户端
	}
}
