package producer_consumer

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	name int
	value int
}

func Producer(wg *sync.WaitGroup, products chan <- Product, name int, stop *bool){
	for !*stop{
		product := Product{
			name:name,
			value: rand.Int(),
		}

		products <- product

		fmt.Printf("producer %v produce a product:%v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

func Consumer(wg *sync.WaitGroup, products <- chan Product, name int){
	for product := range products{
		fmt.Printf("consumer %v consume a product: %v\n", name, product)
		time.Sleep(time.Duration(200+rand.Intn(1000)) * time.Millisecond)
	}
	wg.Done()
}

