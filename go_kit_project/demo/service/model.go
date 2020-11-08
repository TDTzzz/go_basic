package service

//2.定义Model
type DemoReq struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type DemoRes struct {
	IsHealth bool `json:"is_health"`
}
