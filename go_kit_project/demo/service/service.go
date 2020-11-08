package service

import "log"

//1.定义service的接口
type DemoService interface {
	HealthCheck(req DemoReq) DemoRes
}

//2.创建实现此interface的结构体
type DemoServiceImpl struct {
}

func NewService() DemoService {
	return &DemoServiceImpl{}
}

func (d DemoServiceImpl) HealthCheck(req DemoReq) DemoRes {
	log.Println(req.Ip, ":", req.Port, "健健康康！！！")
	return DemoRes{IsHealth: true}
}
