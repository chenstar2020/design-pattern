package singleton

import "sync"

type Singleton interface {
	foo()
}


type singleton struct{}
func (s singleton) foo(){}

var _ Singleton = (*singleton)(nil)
var(
	instance *singleton
	once sync.Once
)

func GetInstance() Singleton{
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
