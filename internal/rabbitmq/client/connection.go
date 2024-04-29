package client

import (
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"go-booking-system/internal/rabbitmq"
)

type RabbitMQConnection struct {
	conn *amqp.Connection
}

func connect(username, password, host, port, vhost string) (*amqp.Connection, error) {
	amqpUrl := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vhost)
	conn, err := amqp.Dial(amqpUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewRabbitMQConnection(username, password, host, port, vhost string) RabbitMQConnection {
	conn, err := connect(username, password, host, port, vhost)
	if err != nil {
		panic(err)
	}
	return RabbitMQConnection{
		conn: conn,
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
