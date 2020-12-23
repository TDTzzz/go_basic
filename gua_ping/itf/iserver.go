package itf

//定义server的接口

type IServer interface {
	//server操作
	Start()
	Stop()
	Serve()
	//添加路由
	AddRouter(msgId uint32, router IRouter)
	//资源操作
	GetConnMgr() IConnManager
	//钩子相关
}
