package command

import (
	"fmt"
)

/******抽象命令接口*******/

// Command 抽象命令接口
type Command interface {
	Execute()
}

/******具体命令实现********/

// StartCommand 开启命令
type StartCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard)*StartCommand{
	return &StartCommand{
		mb: mb,
	}
}

func (s *StartCommand) Execute() {
	s.mb.Start()
}

var _ Command = (*StartCommand)(nil)

// RebootCommand 重启命令
type RebootCommand struct{
	mb *MotherBoard
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand{
	return &RebootCommand{
		mb: mb,
	}
}

func (r *RebootCommand) Execute(){
	r.mb.Reboot()
}

var _ Command = (*RebootCommand)(nil)

// CloseCommand 关闭命令
type CloseCommand struct {
	mb *MotherBoard
}

func NewCloseCommand(mb *MotherBoard) *CloseCommand{
	return &CloseCommand{
		mb: mb,
	}
}

func (c *CloseCommand) Execute(){
	c.mb.Close()
}

// Box 遥控器
type Box struct {
	startButton  Command      //按钮1  开始按钮
	rebootButton Command      //按钮2  重启按钮
	closeButton  Command      //按钮3  关闭按钮
}

func NewBox(startButton, rebootButton, closeButton Command) *Box{
	return &Box{
		startButton: startButton,
		rebootButton: rebootButton,
		closeButton: closeButton,
	}
}

func (b *Box)PressStartButton(){
	b.startButton.Execute()
}

func (b *Box)PressRebootButton(){
	b.rebootButton.Execute()
}

func (b *Box)PressCloseButton(){
	b.closeButton.Execute()
}


// MotherBoard 主板  负责执行具体命令
type MotherBoard struct{}

func (*MotherBoard) Start(){
	fmt.Print("system starting\n")
}

func (*MotherBoard) Reboot(){
	fmt.Print("system rebooting\n")
}

func (*MotherBoard) Close(){
	fmt.Print("system closing\n")
}