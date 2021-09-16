package simpleFactory

import "testing"

func NewAPI(i int) API{
	switch i {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	}
	return nil
}

func TestSimpleFactory(t *testing.T) {
	api := NewAPI(1)
	api.Say("jack")

	api2 := NewAPI(2)
	api2.Say("tom")
}

