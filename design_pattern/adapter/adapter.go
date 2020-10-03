package adapter

//适配者（adapter）的目标接口
type Target interface {
	Request() string
}

//被适配者(adaptee)的目标接口
type Adaptee interface {
	SpecificRequest() string
}

//被适配者(adaptee)的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//被适配者
type adapteeImpl struct{}

//适配器
type adapter struct {
	Adaptee
}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

//
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{adaptee}
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
