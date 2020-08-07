package integration

import (
	"fmt"
	"os"

	"github.com/streadway/amqp"
)

// ConnectionRabbitMQ - connect struct rabbitMQ
type ConnectionRabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

// GetConnectionRabbit - get connection rabbitMQ
func GetConnectionRabbit(rabbitURL string) (ConnectionRabbitMQ, error) {
	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		return ConnectionRabbitMQ{}, err
	}

	ch, err := conn.Channel()
	return ConnectionRabbitMQ{
		Channel: ch,
		Conn:    conn,
	}, err
}

// StartQueeConsumer - handle
func (conn ConnectionRabbitMQ) StartQueeConsumer(queueName, routingKey string, handler func(d amqp.Delivery) bool, concurrency int) error {

	_, err := conn.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = conn.Channel.QueueBind(queueName, routingKey, "telegram", false, nil)
	if err != nil {
		return err
	}

	prefetchCount := concurrency * 4
	err = conn.Channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return err
	}

	msgs, err := conn.Channel.Consume(
		queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for i := 0; i < concurrency; i++ {
		go func() {
			for msg := range msgs {
				if handler(msg) {
					msg.Ack(false)
				} else {
					msg.Nack(false, true)
				}
			}
			fmt.Println("Rabbit consumer closed - critical Error")
			os.Exit(1)
		}()
	}
	return nil
}

// PublishRabbitMQ - publish rabbitMQ queue
func (conn ConnectionRabbitMQ) PublishRabbitMQ(routingKey string, data []byte, exchange string) error {
	return conn.Channel.Publish(
		exchange,
		routingKey,
		true,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}
