package state

import (
	"fmt"
)

type Week interface {
	Today()
	Next(ctx *DayContext)
}

type DayContext struct {
	today Week
}

func NewDayContext() *DayContext{
	return &DayContext{
		today: &Sunday{},
	}
}

func (d *DayContext) Today() {
	d.today.Today()
}

func (d *DayContext) Next() {
	d.today.Next(d)
}

// Sunday 周日
type Sunday struct {}

func (*Sunday) Today(){
	fmt.Printf("today is Sunday\n")
}

func (*Sunday) Next(ctx *DayContext){
	ctx.today = &Monday{}
}

// Monday 周一
type Monday struct {}

func (*Monday) Today(){
	fmt.Printf("today is Monday\n")
}

func (*Monday) Next(ctx *DayContext){
	ctx.today = &TuesDay{}
}

// TuesDay 周二
type TuesDay struct {}

func (*TuesDay) Today(){
	fmt.Printf("today is TuesDay\n")
}

func (*TuesDay) Next(ctx *DayContext){
	ctx.today = &WednesDay{}
}

// WednesDay 周三
type WednesDay struct {}

func (*WednesDay) Today(){
	fmt.Printf("today is WednesDay\n")
}

func (*WednesDay) Next(ctx *DayContext){
	ctx.today = &ThursDay{}
}

// ThursDay 周四
type ThursDay struct {}

func (*ThursDay) Today(){
	fmt.Printf("today is ThursDay\n")
}

func (*ThursDay) Next(ctx *DayContext){
	ctx.today = &FriDay{}
}

// FriDay 周五
type FriDay struct {}

func (*FriDay) Today(){
	fmt.Printf("today is FriDay\n")
}

func (*FriDay) Next(ctx *DayContext){
	ctx.today = &SaturDay{}
}

// SaturDay 周六
type SaturDay struct {}

func (*SaturDay) Today(){
	fmt.Printf("today is SaturDay\n")
}

func (*SaturDay) Next(ctx *DayContext){
	ctx.today = &Sunday{}
}