package shop

import "fmt"

//商店接口
type Shop interface {
	sell()
}

//顾客类
type Customer struct {

}

func (c *Customer) goShopping(s Shop){
	s.sell()
}



type WuyuanShop struct{}

func (w *WuyuanShop) sell(){
	fmt.Println("五元超市购物")
}

type ShaoguanShop struct {
}

func (s *ShaoguanShop)sell(){
	fmt.Println("韶关超市购物")
}
