package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 构建channel, 容量为10
	ch := make(chan int, 10)
	defer close(ch)

	// 生产消息
	go produce(ch)
	// 消费消息
	go consume(ch)

	for {
	}
}

// 生产者, 往channel写数据
func produce(ch chan<- int) {
	for {
		time.Sleep(time.Second)
		num := rand.Intn(10)
		ch <- num
		fmt.Printf("produce: %d\n", num)
	}
}

// 消费者, 从channel读数据
func consume(ch <-chan int) {
	for {
		time.Sleep(time.Second)
		num := <-ch
		fmt.Printf("consume: %d\n", num)
	}
}
