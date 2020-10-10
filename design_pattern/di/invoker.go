package di

type Invoker interface {
	//init cfg returns parse cfg error
	InitCfg(cfg []byte, cfgType string) error
	//InitCaller returns init caller error
	Run() error
}

type InvokerFunc func() Invoker
