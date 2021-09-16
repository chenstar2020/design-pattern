package adapter

// Target 适配的目标接口 提供给外部使用
type Target interface {
	Request() string
}

// Adaptee 被适配的目标接口  内部的具体实现
type Adaptee interface {
	SpecificRequest() string
}

func NewAdaptee() Adaptee{
	return &adapteeImp1{}
}


//被适配的目标类  被适配类的具体实现
type adapteeImp1 struct{}

func (a adapteeImp1) SpecificRequest() string {
	return "adaptee1 method"
}

var _ Adaptee = (*adapteeImp1)(nil)


func NewAdapter(adaptee Adaptee) Target{
	return &adapter{
		Adaptee: adaptee,
	}
}

//适配的目标类  适配类的具体实现  提供给外部调用
type adapter struct {
	Adaptee
}

func (a *adapter) Request() string{
	return a.SpecificRequest()
}

var _ Target = (*adapter)(nil)