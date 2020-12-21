package zitf

//定义server的接口

type IServer interface {
	//server操作
	Start()
	Stop()
	Serve()
	//添加路由
	//AddRouter()
	//钩子相关
}
