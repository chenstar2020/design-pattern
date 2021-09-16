package abstractFactory

import "fmt"

// OrderMainDAO 订单主记录
type OrderMainDAO interface {
	SaveOrderMain()
}

// OrderDetailDAO 订单详细记录
type OrderDetailDAO interface {
	SaveOrderDetail()
}

// DAOFactory 抽象模式工厂接口
type DAOFactory interface {
	CreateOrderMainDAO() OrderMainDAO
	CreateOrderDetailDAO() OrderDetailDAO
}

// RDBMainDAO 关系型数据库OrderMainDAO实现
type RDBMainDAO struct{}

func (R RDBMainDAO) SaveOrderMain() {
	fmt.Print("rdb main save\n")
}

var _ OrderMainDAO = (*RDBMainDAO)(nil)

// RDBDetailDAO 关系型数据库OrderDetailDAO实现
type RDBDetailDAO struct{}

func (R RDBDetailDAO) SaveOrderDetail() {
	fmt.Print("rdb detail save\n")
}

var _ OrderDetailDAO = (*RDBDetailDAO)(nil)

// RDBDAOFactory RDB 抽象工厂实现
type RDBDAOFactory struct{}

func (R RDBDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &RDBMainDAO{}
}

func (R RDBDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &RDBDetailDAO{}
}

var _ DAOFactory = (*RDBDAOFactory)(nil)


// XMLMainDAO xml main存储
type XMLMainDAO struct{}

func (X XMLMainDAO) SaveOrderMain() {
	fmt.Print("xml main save\n")
}

var _ OrderMainDAO = (*XMLMainDAO)(nil)

// XMLDetailDAO xml detail存储
type XMLDetailDAO struct{}

func (X XMLDetailDAO) SaveOrderDetail() {
	fmt.Print("xml detail save\n")
}

var _ OrderDetailDAO = (*XMLDetailDAO)(nil)

type XMLDAOFactory struct{}

func (X XMLDAOFactory) CreateOrderMainDAO() OrderMainDAO {
	return &XMLMainDAO{}
}

func (X XMLDAOFactory) CreateOrderDetailDAO() OrderDetailDAO {
	return &XMLDetailDAO{}
}

var _ DAOFactory = (*XMLDAOFactory)(nil)

