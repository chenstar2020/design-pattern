package observer

import (
	"fmt"
)

type Subject struct {
	observers []Observer
	context   string
}

func NewSubject() *Subject{
	return &Subject{
		observers: []Observer{},
	}
}

// Attach 添加观察者
func (s *Subject) Attach(o Observer){
	s.observers = append(s.observers, o)
}

// Deletion 删除观察者
func (s *Subject) Deletion(o Observer){

}

func (s *Subject) notify() {
	for _, o := range s.observers{
		o.Update(s)
	}
}

func (s *Subject) UpdateContext(context string){
	s.context = context
	s.notify()
}

type Observer interface {
	Update(*Subject)
}

type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{
		name: name,
	}
}

func (r *Reader) Update(s *Subject){
	fmt.Printf("%s receive %s\n", r.name, s.context)
}