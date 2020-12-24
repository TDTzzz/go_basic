package net

import "go_basic/gua_ping/itf"

type Request struct {
	conn itf.IConnection
	msg  itf.IMessage
}

func (r *Request) GetConnection() itf.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
