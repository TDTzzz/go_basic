package main

import (
	"fmt"
	net2 "go_basic/gua_ping/net"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("Client Test ... start")

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err,exit!")
		return
	}

	for {
		//发封包消息
		dp := net2.NewDataPack()
		msg, _ := dp.Pack(net2.NewMsgPackage(0, []byte("Gua Ping Client Test Message!!!!!!")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("write error err", err)
			return
		}
		//先读head
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData)
		if err != nil {
			fmt.Println("read head error",err)
			break
		}

		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		if msgHead.GetDataLen() > 0 {
			msg := msgHead.(*net2.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
		time.Sleep(1 * time.Second)
	}
}
