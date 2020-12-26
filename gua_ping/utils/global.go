package utils

import "go_basic/gua_ping/itf"

type Global struct {
	//Server
	TcpServer itf.IServer
	Host      string //当前服务器的主机
	TcpPort   int    //当前服务器主机监听端口号
	Name      string //当前服务器名称
	//
	WorkerPoolSize   uint32
	MaxPacketSize    uint32
	MaxMsgChanLen    uint32 //SendMsgMsg发送消息的缓冲最大长度
	MaxWorkerTaskLen uint32 //业务工作Worker对应负责的任务队列最大任务存储数量
}

//定义一个全局对象
var GlobalObj *Global

func (g *Global) Reload() {

}

//初始化全局变量，设定一些默认值
func init() {
	GlobalObj = &Global{
		Name:             "GuaPingServerApp",
		Host:             "0.0.0.0",
		TcpPort:          8999,
		WorkerPoolSize:   10,
		MaxPacketSize:    4096,
		MaxMsgChanLen:    1024,
		MaxWorkerTaskLen: 1024,
	}
	
	//从配置文件里加载一些用户配置
	GlobalObj.Reload()
}
