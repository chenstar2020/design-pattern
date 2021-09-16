package command

import "testing"

func TestCommand(t *testing.T) {
	mb := &MotherBoard{}
	startCommand := NewStartCommand(mb)
	rebootCommand := NewRebootCommand(mb)
	closeCommand := NewCloseCommand(mb)

	box := NewBox(startCommand, rebootCommand, closeCommand)
	box.PressStartButton()       //按下按钮1
	box.PressRebootButton()      //按下按钮2
	box.PressCloseButton()       //按下按钮3
}