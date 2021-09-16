package factoryMethod

type Operator interface {
	SetA(int)
	SetB(int)
	Result() int
}

type OperatorBase struct {
	a, b int
}

func (o *OperatorBase) SetA(a int){
	o.a = a
}

func (o *OperatorBase) SetB(b int){
	o.b = b
}


// OperatorFactory 抽象工厂接口
type OperatorFactory interface {
	Create() Operator
}

// PlusOperatorFactory 加法工厂实现
type PlusOperatorFactory struct{}

func (p PlusOperatorFactory) Create() Operator {
	return &PlusOperator{
		&OperatorBase{},
	}
}

var _ OperatorFactory = (*PlusOperatorFactory)(nil)

// MinusOperatorFactory 减法工厂实现
type MinusOperatorFactory struct{}

func (m MinusOperatorFactory) Create() Operator {
	return &MinusOperator{
		&OperatorBase{},
	}
}

var _ OperatorFactory = (*MinusOperatorFactory)(nil)


// PlusOperator 加法实现
type PlusOperator struct {
	*OperatorBase
}

func (p *PlusOperator) Result() int {
	return p.a + p.b
}

var _ Operator = (*PlusOperator)(nil)


// MinusOperator 减法实现
type MinusOperator struct {
	*OperatorBase
}

func (m *MinusOperator) Result() int {
	return m.a - m.b
}

var _ Operator = (*MinusOperator)(nil)