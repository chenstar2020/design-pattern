package main

import (
	"./producer_consumer"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)



func main(){
/*	subject := observer.NewSubject()

	reader1 := observer.NewReader("reader1")
	reader2 := observer.NewReader("reader2")

	subject.Attach(reader1)
	subject.Attach(reader2)

	subject.UpdateContext("observer mode")*/

	var wgp sync.WaitGroup
	var wgc sync.WaitGroup

	stop := false
	products := make(chan producer_consumer.Product, 10)
	for i := 0; i < 5; i++ {
		go producer_consumer.Producer(&wgp, products, i, &stop)
		go producer_consumer.Consumer(&wgc, products, i)
		wgp.Add(1)
		wgc.Add(1)
	}

	time.Sleep(time.Duration(1) * time.Second)
	stop = true
	wgp.Wait()
	close(products)
	wgc.Wait()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	err := r.Run("127.0.0.1:8765")
	if err != nil {
		panic(err)
	}
}
