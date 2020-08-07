package integration

import (
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

// PublishRabbitMQ - publish rabbitMQ queue
func (conn ConnectionRabbitMQ) PublishRabbitMQ(routingKey string, data []byte, exchange string) error {
	return conn.Channel.Publish(
		// exchange - yours may be different
		exchange,
		routingKey,
		// mandatory - we don't care if there I no queue
		false,
		// immediate - we don't care if there is no consumer on the queue
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         data,
			DeliveryMode: amqp.Persistent,
		})
}
