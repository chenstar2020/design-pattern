package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	component := ConcreteComponent{}
	addDecorator := WarpAddDecorator(component, 10)
	result := addDecorator.Calc()
	if result != 10 {
		t.Fatal("test AddDecorator fail")
	}
	mulDecorator := WarpMulDecorator(component, 8)
	result = mulDecorator.Calc()
	if result != 0 {
		t.Fatal("test MulDecorator fail")
	}
}