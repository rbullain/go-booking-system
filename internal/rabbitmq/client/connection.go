package client

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-booking-system/cmd/rabbitmq/config"
	"go-booking-system/internal/rabbitmq"
)

var queuesNames = []string{
	config.QueueUserCreated, config.QueueReservationCreated,
}

type IMessagingClient interface {
	PublishOnQueue(payload rabbitmq.IRabbitMQPayload, queueName string) error
	Subscribe(consumerName string, handler func(amqp.Delivery)) error
	Close()
}

type BookingMessageClient struct {
	conn *amqp.Connection
}

var _ IMessagingClient = BookingMessageClient{}

func (connection BookingMessageClient) InitializeQueues() {
	ch, err := connection.conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		config.Exchange,
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for _, queueName := range queuesNames {
		queue, err := ch.QueueDeclare(
			queueName,
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			panic(err)
		}

		err = ch.QueueBind(
			queue.Name,
			queue.Name,
			config.Exchange,
			false,
			nil,
		)
		if err != nil {
			panic(err)
		}
	}
}

func NewRabbitMQConnection(username, password, host, port, vhost string) BookingMessageClient {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vhost)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		panic(err)
	}
	return BookingMessageClient{
		conn: conn,
	}
}

func (connection BookingMessageClient) Close() {
	if connection.conn != nil {
		connection.conn.Close()
	}
}

func (connection BookingMessageClient) publish(body []byte, queueName string) error {
	ch, err := connection.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	err = ch.PublishWithContext(
		context.TODO(),
		config.Exchange,
		queueName,
		false,
		false,
		amqp.Publishing{
			Body:        body,
			ContentType: "application/json",
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func (connection BookingMessageClient) PublishOnQueue(payload rabbitmq.IRabbitMQPayload, queueName string) error {
	body, err := json.Marshal(payload)
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

func (connection BookingMessageClient) Subscribe(queueName string, handler func(amqp.Delivery)) error {
	ch, err := connection.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
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
