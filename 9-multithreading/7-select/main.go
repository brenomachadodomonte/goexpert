package main

import (
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0
	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{ID: i, Msg: "Hello from RabbitMQ"}
			c1 <- msg
		}
	}()

	go func() {
		for {
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 2)
			msg := Message{ID: i, Msg: "Hello from Kafka"}
			c2 <- msg
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			println("Received", msg1.Msg, "ID", msg1.ID)
		case msg2 := <-c2:
			println("Received", msg2.Msg, "ID", msg2.ID)
		case <-time.After(time.Second * 3):
			println("Timout")
			//default:
			//	println("default")
		}
	}
}
