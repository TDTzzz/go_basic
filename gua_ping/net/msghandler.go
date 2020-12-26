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

func (mh *MsgHandler) DoMsgHandler(request itf.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
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

//启动worker工作池
func (mh *MsgHandler) StartWorkerPool() {
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		mh.TaskQueue[i] = make(chan itf.IRequest, utils.GlobalObj.MaxWorkerTaskLen)
		go mh.StartOnWorker(i, mh.TaskQueue[i])
	}
}

//启动一个Worker工作流程
func (mh *MsgHandler) StartOnWorker(workerID int, taskQueue chan itf.IRequest) {
	fmt.Println("Worker ID = ", workerID, " is started.")

	for {
		select {
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

func (mh *MsgHandler) SendMsgToTaskQueue(request itf.IRequest) {
	//得到需处理此条连接的WorkerID
	workerID := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	//将请求消息发给任务队列
	mh.TaskQueue[workerID] <- request
}

func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		Apis:           make(map[uint32]itf.IRouter),
		WorkerPoolSize: utils.GlobalObj.WorkerPoolSize,
		TaskQueue:      make([]chan itf.IRequest, utils.GlobalObj.WorkerPoolSize),
	}
}
