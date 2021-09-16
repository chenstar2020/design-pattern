package bridge

import "fmt"

type AbstractMessage interface {
	SendMessage(text, to string)
}

type MessageImplementer interface {
	Send(text, to string)
}

type MessageSMS struct{}

func (m MessageSMS) Send(text, to string) {
	fmt.Printf("发送短信：%v To：%v\n", text, to)
}

var _ MessageImplementer = (*MessageSMS)(nil)

func ViaSMS() MessageImplementer{
	return &MessageSMS{}
}

type MessageEmail struct{}

func (m MessageEmail) Send(text, to string) {
	fmt.Printf("发送Email：%v To：%v\n", text, to)
}

var _ MessageImplementer = (*MessageEmail)(nil)

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

type CommonMessage struct {         //使用关联 将一个类对象嵌入到另一个类对象中
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage{
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string){
	m.method.Send(fmt.Sprintf("[Common] %s", text), to)
}

type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage{
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string){
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}





