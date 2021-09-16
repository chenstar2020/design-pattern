package mediator

import "testing"

func TestMediator(t *testing.T){
	//先创建中介者
	m 		:= GetMediatorInstance()
	m.CD 	= &CDDriver{}
	m.CPU 	= &CPU{}
	m.Sound = &SoundCard{}
	m.Video = &VideoCard{}
	m.HeadSet = &HeadSet{}
	m.Monitor = &Monitor{}

	m.CD.ReadData()
}