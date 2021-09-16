package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do()string{
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p *Proxy) Do()string{
	var res string

	//在调用真实对象之前的工作，检查缓存，判断权限，实例化真实对象等
	res += "pre:"

	res += p.real.Do()

	res += ":after"
	//调用之后的操作，缓存结果，对结果进行处理等
	return res
}