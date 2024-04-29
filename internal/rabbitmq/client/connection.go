package client

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-booking-system/internal/rabbitmq"
)

type IMessagingClient interface {
	Connect(username, password, host, port, vhost string)
	PublishOnQueue(payload rabbitmq.IPayload, queueName string) error
	SubscribeToQueue(queueName, consumerName string, handler func(msg amqp.Delivery)) error
	Close()
}

type RabbitMQConnection struct {
	conn *amqp.Connection
}

var _ IMessagingClient = RabbitMQConnection{}

func (connection RabbitMQConnection) Connect(username, password, host, port, vhost string) {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vhost)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		panic(err)
	}
	connection.conn = conn
}

func NewRabbitMQConnection(username, password, host, port, vhost string) RabbitMQConnection {
	conn := RabbitMQConnection{}
	conn.Connect(username, password, host, port, vhost)
	return conn
}

func (connection RabbitMQConnection) Close() {
	if connection.conn != nil {
		connection.conn.Close()
	}
}

func (connection RabbitMQConnection) publish(body []byte, queueName string) error {
	ch, err := connection.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.PublishWithContext(
		context.TODO(),
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			Body: body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (connection RabbitMQConnection) PublishOnQueue(payload rabbitmq.IPayload, queueName string) error {
	body, err := payload.Encode()
	if err != nil {
		return err
	}

	err = connection.publish(body, queueName)
	if err != nil {
		return err
	}

	return nil
}

func consume(msgs <-chan amqp.Delivery, handler func(msg amqp.Delivery)) {
	for msg := range msgs {
		handler(msg)
	}
}

func (connection RabbitMQConnection) SubscribeToQueue(queueName, consumerName string, handler func(msg amqp.Delivery)) error {
	ch, err := connection.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue.Name,
		consumerName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go consume(msgs, handler)
	return nil
}
