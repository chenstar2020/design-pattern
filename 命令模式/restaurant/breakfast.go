package restaurant

import "fmt"

/********抽象命令接口**********/

// Breakfast 早餐类接口
type Breakfast interface {
	cooking()
}

/********具体命令类************/

// ChangFen 肠粉子类
type ChangFen struct{
	c *Chief
}

func (c ChangFen) cooking() {
	c.c.doChangFen()
}

var _ Breakfast = (*ChangFen)(nil)

// HunTun 馄饨子类
type HunTun struct{
	c *Chief
}

func (h HunTun) cooking() {
	h.c.doHunTun()
}

var _ Breakfast = (*HunTun)(nil)

/*********命令接收者*************/

type Chief struct {

}

func (c *Chief) doChangFen(){
	fmt.Println("开始做肠粉")
}

func (c *Chief) doHunTun(){
	fmt.Println("开始做馄饨")
}