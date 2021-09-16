package strategy

import "testing"

func TestPayment_Pay(t *testing.T) {
	cash := &Cash{}
	cashPayment := NewPayment("招商银行", "12345", "$45,000", cash)
	cashPayment.Pay()

	card := &Card{}
	cashPayment = NewPayment("招商银行", "12345", "$45,000", card)
	cashPayment.Pay()
}