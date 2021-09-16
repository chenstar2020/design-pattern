package pubsub

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}   		//订阅者为一个管道
	topicFunc func(v interface{}) bool  //主题为一个过滤器
)

// Publisher 发布者对象
type Publisher struct {
	m		sync.RWMutex
	buffer  int
	timeout time.Duration
	subscribers map[subscriber]topicFunc
}

func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher{
	return &Publisher{
		buffer: buffer,
		timeout: publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// Subscribe 添加一个新的订阅者，订阅全部主题
func (p *Publisher)Subscribe() chan interface{}{
	return p.SubscriberTopic(nil)
}

// SubscriberTopic 添加一个新的订阅者
func (p *Publisher) SubscriberTopic(topic topicFunc) chan interface{}{
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// Evict 退出订阅
func (p *Publisher) Evict(sub chan interface{}){
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribers, sub)
	close(sub)
}

// Publish 发布一个主题
func (p *Publisher) Publish(v interface{}){
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribers{
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup){
	defer wg.Done()

	if topic != nil && !topic(v){    //主题过滤器
		return
	}

	select{
	case sub <- v:
	case <- time.After(p.timeout):
	}
}

// Close 关闭发布者对象
func (p *Publisher) Close(){
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribers{
		delete(p.subscribers, sub)
		close(sub)
	}
}



