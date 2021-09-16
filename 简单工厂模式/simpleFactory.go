package simpleFactory

import "fmt"

type API interface {
	Say(name string)
}


type hiAPI struct{}

func (h hiAPI) Say(name string)  {
	fmt.Println("hi,", name)
}

var _ API = (*hiAPI)(nil)


type helloAPI struct{}

func (h helloAPI) Say(name string)  {
	fmt.Println("hello,", name)
}

var _ API = (*helloAPI)(nil)