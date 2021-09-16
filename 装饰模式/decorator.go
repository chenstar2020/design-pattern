package decorator

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (c ConcreteComponent) Calc() int {
	return 0
}

var _ Component = (*ConcreteComponent)(nil)


type MulDecorator struct {
	Component                 //将接口嵌入自身  由自身对象决定扩展 嵌入的接口对象
	num int
}

func WarpMulDecorator(c Component, num int)Component{
	return &MulDecorator{
		Component: c,
		num: num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func (a AddDecorator) Calc() int {
	return a.Component.Calc() + a.num
}

var _ Component = (*AddDecorator)(nil)

func WarpAddDecorator(c Component, num int) Component{
	return &AddDecorator{
		Component: c,
		num: num,
	}
}




