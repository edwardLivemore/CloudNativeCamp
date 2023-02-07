package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
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

// 获取Goroutine的ID
func GetGid() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

// 生产者, 往channel写数据
func produce(ch chan<- int) {
	for {
		time.Sleep(time.Second)
		num := rand.Intn(10)
		ch <- num
		fmt.Printf("Goroutine id: %v, produce: %d\n", GetGid(), num)
	}
}

// 消费者, 从channel读数据
func consume(ch <-chan int) {
	for {
		time.Sleep(time.Second)
		num := <-ch
		fmt.Printf("Goroutine id: %v, consume: %d\n", GetGid(), num)
	}
}
