package znet

import "fmt"



var logo
//Server的接口实现
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s Server) Start() {
	panic("implement me")
}

func (s Server) Stop() {
	panic("implement me")
}

func (s Server) Serve() {
	panic("implement me")
}




func init()  {
	fmt.Println("")
}



