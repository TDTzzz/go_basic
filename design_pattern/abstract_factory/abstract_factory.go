package abstract_factory

import "fmt"

//订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

//订单详情
type OrderDetailDAO interface {
	SaveOrderDetail()
}

//抽象模式机工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

type RDBMainDAO struct{}

func (*RDBMainDAO) SaveOrderMain() {
	fmt.Println("rdb main save")
}

type RDBDetailDAO struct{}

func (*RDBDetailDAO) SaveOrderDetail() {
	fmt.Println("rdb detail save")
}

//RDB抽象工厂的实现
type RDBDAOFactory struct{}

func (*RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (*RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

//XMLMainDAO XML存储
type XMLMainDao struct{}

func (*XMLMainDao) SaveOrderMain() {
	fmt.Println("xml main save")
}

type XMLDetailDao struct{}

func (*XMLDetailDao) SaveOrderDetail() {
	fmt.Println("xml detail save")
}

type XMLDAOFactory struct{}

func (*XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDao{}
}

func (*XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDao{}
}
