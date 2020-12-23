package net

import (
	"context"
	"errors"
	"fmt"
	"go_basic/gua_ping/itf"
	"go_basic/gua_ping/utils"
	"io"
	"net"
	"sync"
)

type Connection struct {
	TcpServer  itf.IServer
	Conn       *net.TCPConn
	ConnID     uint32
	MsgHandler itf.IMsgHandler

	ctx context.Context
	sync.RWMutex

	//消息管道
	msgChan     chan []byte
	msgBuffChan chan []byte
	//当前连接的关闭状态
	isClosed bool
}

func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	c.RLock()
	if c.isClosed {
		c.RUnlock()
		return errors.New("connection closed when send msg")
	}
	c.RUnlock()

	//封包并发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}
	//写回客户端
	c.msgChan <- msg
	return nil
}

func (c *Connection) SendBuffMsg(msgId uint32, data []byte) error {
	c.RLock()
	if c.isClosed == true {
		c.RUnlock()
		return errors.New("Connection closed when send buff msg")
	}
	c.RUnlock()
	//封包并发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("Pack error msg ")
	}
	c.msgBuffChan <- msg
	return nil
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Start() {

	go c.StartReader()
	go c.StartWriter()

	//执行钩子
}

func (c *Connection) StartReader() {
	//启动connection
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println(c.RemoteAddr().String(), "[conn Reader exit!]")
	defer c.Stop()

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			dp := NewDataPack()
			//读Client的Msg head
			headData := make([]byte, dp.GetHeadLen())
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				fmt.Println("read msg head error ", err)
				return
			}

			//拆包得到msgid和dataLen
			msg, err := dp.Unpack(headData)
			if err != nil {
				fmt.Println("unpack error", err)
				return
			}
			//根据dataLen读取data，放在msg.Data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					fmt.Println("read msg data error", err)
					return
				}
			}
			msg.SetData(data)

			//将msg包装到req里，便于分发
			req := Request{
				conn: c,
				msg:  msg,
			}

			if utils.GlobalObj.WorkerPoolSize > 0 {
				//交给worker处理,非阻塞

			} else {
				go c.MsgHandler.DoMsgHandler(&req)
			}

		}
	}
}

func (c *Connection) StartWriter() {

}

func (c *Connection) Stop() {

	fmt.Println("Conn Stop()...ConnID = ", c.ConnID)

	c.Lock()
	defer c.Unlock()

	//如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用
	//c.TcpServer.CallOnConnStop(c)

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// 关闭socket链接
	c.Conn.Close()
	//关闭Writer
	//c.cancel()

	//将链接从连接管理器中删除
	c.TcpServer.GetConnMgr().Remove(c)
	//关闭该链接全部管道
	close(c.msgBuffChan)
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func NewConnection(server itf.IServer, conn *net.TCPConn, connID uint32, msgHandler itf.IMsgHandler) *Connection {
	c := &Connection{
		TcpServer:  server,
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
	}
	//将这个新建的Conn绑定到链接管理中
	c.TcpServer.GetConnMgr().Add(c)
	return c
}
