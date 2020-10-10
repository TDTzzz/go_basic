package di

//参考beemod的依赖注入写法

type BeeMod struct {
	invokers []Invoker
	cfgByte  []byte

}
