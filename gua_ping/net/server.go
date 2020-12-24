package net

import (
	"fmt"
	"go_basic/gua_ping/itf"
	"go_basic/gua_ping/utils"
	"net"
)

var guaPingLogo = `
                                
 ████████   ██     ██          █
 ██         ██     ██        █   █
 ██    ██   ██     ██       ███████  
 ██     █   ██     ██     ████   ████ 
 ████████   █████████    ████     ████

`

/*
   创建一个服务器句柄
*/
func NewServer() itf.IServer {
	s := &Server{
		Name:       utils.GlobalObj.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObj.Host,
		Port:       utils.GlobalObj.TcpPort,
		ConnMgr:    NewConnManager(),
		msgHandler: NewMsgHandler(),
	}
	return s
}

//Server的接口实现
type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	ConnMgr    itf.IConnManager
	msgHandler itf.IMsgHandler
	//钩子函数
	OnConnStart func(conn itf.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn itf.IConnection)
}

func (s *Server) CallOnConnStart(conn itf.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

func (s *Server) CallOnConnStop(conn itf.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

func (s *Server) GetConnMgr() itf.IConnManager {
	return s.ConnMgr
}

func (s *Server) AddRouter(msgId uint32, router itf.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}

//开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server name %s,listener at IP:%s,Port %d is starting\n", s.Name, s.IP, s.Port)

	go func() {
		//1.获得一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		//2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start gua_ping server  ", s.Name, " succ, now listening...")

		var cid uint32
		cid = 0
		//启动server网络连接业务
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())
			//todo 设置了服务器最大连接控制数，如果超过则关闭此新连接

			//todo 处理新连接请求的业务方法，handler和conn绑定
			dealConn := NewConnection(s, conn, cid, s.msgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	panic("implement me")
}

//运行服务
func (s *Server) Serve() {
	s.Start()

	select {} //阻塞
}

func init() {
	fmt.Println(guaPingLogo)
}
