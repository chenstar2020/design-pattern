package strategy

import "fmt"

type Payment struct {
	context *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money string
}


func NewPayment(name, cardId string, money string, strategy PaymentStrategy) *Payment{
	return &Payment{
		context: &PaymentContext{
			Name: name,
			CardID: cardId,
			Money: money,
		},
		strategy: strategy,
	}
}

func (p *Payment)Pay(){
	p.strategy.Pay(p.context)
}

type PaymentStrategy interface {
	Pay(ctx *PaymentContext)
}

type Cash struct{}

func (c *Cash) Pay(pc *PaymentContext){
	fmt.Println("现金支付：", pc.Money)
}

type Card struct {}

func (c *Card) Pay(pc *PaymentContext){
	fmt.Println("刷卡支付：", pc.CardID, pc.Name)
}

