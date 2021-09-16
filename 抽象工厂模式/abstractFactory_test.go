package abstractFactory

import "testing"

func TestAbstractFactory(t *testing.T){
	rdbFactory := RDBDAOFactory{}
	mainDAO := rdbFactory.CreateOrderMainDAO()
	detailDAO := rdbFactory.CreateOrderDetailDAO()

	mainDAO.SaveOrderMain()
	detailDAO.SaveOrderDetail()

	xmlFactory := XMLDAOFactory{}
	mainDAO = xmlFactory.CreateOrderMainDAO()
	detailDAO = xmlFactory.CreateOrderDetailDAO()
	mainDAO.SaveOrderMain()
	detailDAO.SaveOrderDetail()
}
