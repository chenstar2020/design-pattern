package shop

import "testing"

func TestShop(t *testing.T){
	w := &WuyuanShop{}
	s := &ShaoguanShop{}

	customer := &Customer{}
	customer.goShopping(s)
	customer.goShopping(w)
}
