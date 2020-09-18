package svc

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello," + request
	return nil
}
