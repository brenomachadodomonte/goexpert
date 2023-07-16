package main

import "github.com/brenomachadodomonte/goexpert/events/pkg/rabbitmq"

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = rabbitmq.Publish(ch, "amq.direct", "Hello World")
	if err != nil {
		panic(err)
	}
}
