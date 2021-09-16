package main

import (
	"./pubsub"
	"fmt"
	"strings"
	"time"
)

func main(){
	p := pubsub.NewPublisher(100 * time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()              //接收所有消息
	golang := p.SubscriberTopic(func (v interface{}) bool{    //添加一个新的订阅者 返回一个通信管道
		if s, ok := v.(string); ok{
			fmt.Println("订阅者golang接收到消息：", s)
			return strings.Contains(s, "golang")
		}
		return false
	})
	world := p.SubscriberTopic(func (v interface{}) bool{
		if s, ok := v.(string); ok{
			fmt.Println("订阅者world接收到消息：", s)
			return strings.Contains(s, "world")
		}
		return false
	})

	p.Publish("hello, world!")
	p.Publish("hello, golang!")

	world1 := p.SubscriberTopic(func (v interface{}) bool{
		if s, ok := v.(string); ok{
			fmt.Println("订阅者world1接收到消息：", s)
			return strings.Contains(s, "world")
		}
		return false
	})
	fmt.Println("睡眠一秒钟。。。。")
	time.Sleep(time.Second)

	p.Publish("fuck world")
	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()
	go func() {
		for msg := range golang{
			fmt.Println("golang: ", msg)
		}
	}()
	go func() {
		for msg := range world{
			fmt.Println("world: ", msg)
		}
	}()
	go func() {
		time.Sleep(500 * time.Millisecond)
		for msg := range world1{
			fmt.Println("world1: ", msg)
		}
	}()
	time.Sleep(3 * time.Second)
}
