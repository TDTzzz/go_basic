package itf

import "net"

type IConnection interface {
	//操作
	Start()
	Stop()
	//资源操作
	GetConnID() uint32
	RemoteAddr() net.Addr

	//直接将msg数据发给远程TCP客户端
	SendMsg(msgId uint32, data []byte) error
	SendBuffMsg(msgId uint32, data []byte) error
}
