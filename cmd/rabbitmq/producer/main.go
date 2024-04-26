package main

import (
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	//
	//channel, err := conn.Channel()
	//failOnError(err, "Failed to open a channel")
	//
	//// Exchange
	//err = channel.ExchangeDeclare(config.Exchange, "topic", true, false, false, false, nil)
	//failOnError(err, "Failed to declare an exchange")
	//
	//// Queues and binds
	//// _, err := channel.QueueDeclare(queue, true, false, false, false, nil)
	//// failOnError(err, fmt.Sprintf("Failed to declare the %s queue", queue))
	////
	//// err = channel.QueueBind(queue, queue+".*", exchange, false, nil)
	//// failOnError(err, fmt.Sprintf("Failed to bind the %s queue", queue))
}
