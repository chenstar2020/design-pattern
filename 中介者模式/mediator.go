package mediator

import (
	"fmt"
	"strings"
)
/*
 *中介者模式：简化对象之间的关联性 使用中介和别的对象通信
 *在以下例子中，CD驱动使用中介调用CPU CPU再使用中介调用Video和Sound
 */

// CDDriver CD驱动
type CDDriver struct {
	Data string
}

func (c *CDDriver) ReadData(){
	c.Data = "music,image"

	fmt.Printf("CDDriver: reading data %s\n", c.Data)
	GetMediatorInstance().changed(c)
}

type CPU struct {
	Video string
	Sound string
}

func (c *CPU) Process(data string){
	sp := strings.Split(data, ",")
	c.Sound = sp[0]
	c.Video = sp[1]

	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.Sound, c.Video)
	GetMediatorInstance().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string){
	v.Data = data
	fmt.Printf("VideoCard: display %s\n", v.Data)
	GetMediatorInstance().changed(v)
}


type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string){
	s.Data = data
	fmt.Printf("SoundCard: play %s\n", s.Data)
	GetMediatorInstance().changed(s)
}

type HeadSet struct {
	Data string
}

func (s *HeadSet) BroadCast(data string){
	s.Data = data
	fmt.Printf("HeadSet: broadcast %s\n", s.Data)
}

type Monitor struct {
	Data string
}

func (m *Monitor) Movie(data string){
	m.Data = data
	fmt.Printf("Monitor: movie %s\n", m.Data)
}

type Mediator struct {
	CD *CDDriver
	CPU *CPU
	Video *VideoCard
	Sound *SoundCard
	HeadSet *HeadSet
	Monitor *Monitor
}


//中介者
var mediator *Mediator

func GetMediatorInstance() *Mediator{
	if mediator == nil{
		mediator = &Mediator{}
	}
	return mediator
}

func (m *Mediator) changed(i interface{}){
	switch inst := i.(type){
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.Sound.Play(inst.Sound)
		m.Video.Display(inst.Video)
	case *SoundCard:
		m.HeadSet.BroadCast(inst.Data)
	case *VideoCard:
		m.Monitor.Movie(inst.Data)
	}
}