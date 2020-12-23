package net

import (
	"fmt"
	"go_basic/gua_ping/itf"
	"go_basic/gua_ping/utils"
	"strconv"
)

type MsgHandler struct {
	Apis map[uint32]itf.IRouter
	//业务工作Worker池的数量
	WorkerPoolSize uint32
	//任务队列
	TaskQueue []chan itf.IRequest
}

func (m MsgHandler) DoMsgHandler(request itf.IRequest) {
	handler, ok := m.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgID(), " is not FOUND!")
		return
	}

	//执行对应处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

func (mh *MsgHandler) AddRouter(msgId uint32, router itf.IRouter) {
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + strconv.Itoa(int(msgId)))
	}
	//添加msgId和Router的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}

func (m *MsgHandler) StartWorkerPool() {
	panic("implement me")
}

func (m *MsgHandler) SendMsgToTaskQueue(request itf.IRequest) {
	panic("implement me")
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]itf.IRouter),
		WorkerPoolSize: utils.GlobalObj.WorkerPoolSize,
		TaskQueue:      make([]chan itf.IRequest, utils.GlobalObj.WorkerPoolSize),
	}
}
