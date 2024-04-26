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
}
