package net

import "go_basic/gua_ping/itf"

type Request struct {
	conn itf.IConnection
	msg  itf.IMessage
}

func (r *Request) GetConnection() itf.IConnection {
	panic("implement me")
}

func (r *Request) GetData() []byte {
	panic("implement me")
}

func (r *Request) GetMsgID() uint32 {
	panic("implement me")
}
