package facade

import (
	"fmt"
)
func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

// API 统一对外的接口
type API interface {
	Test() string
}

// 统一对外接口的实现
type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string{
	aRet := a.a.TestA()
	bRet := a.b.TestB()

	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

var _ API = (*apiImpl)(nil)


// AModuleAPI A模块接口
type AModuleAPI interface {
	TestA() string
}

func NewAModuleAPI() AModuleAPI{
	return &aModuleImpl{}
}

//A模块接口的实现
type aModuleImpl struct{}

func (a aModuleImpl) TestA() string {
	return "A module running"
}

var _ AModuleAPI = (*aModuleImpl)(nil)


// BModuleAPI B模块接口
type BModuleAPI interface {
	TestB() string
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

//B模块接口的实现
type bModuleImpl struct{}

func (b bModuleImpl) TestB() string {
	return "B module running"
}

var _ BModuleAPI = (*bModuleImpl)(nil)






