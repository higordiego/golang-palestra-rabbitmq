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

// StartConsumer - handle
func (conn ConnectionRabbitMQ) StartQueeConsumer(queueName, routingKey string, handler func(d amqp.Delivery) bool, concurrency int) error {

	// create the queue if it doesn't already exist
	_, err := conn.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}

	// bind the queue to the routing key
	err = conn.Channel.QueueBind(queueName, routingKey, "whatsapp", false, nil)
	if err != nil {
		return err
	}

	// prefetch 4x as many messages as we can handle at once
	prefetchCount := concurrency * 4
	err = conn.Channel.Qos(prefetchCount, 0, false)
	if err != nil {
		return err
	}

	msgs, err := conn.Channel.Consume(
		queueName, // queue
		"",        // consumer
		false,     // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	// create a goroutine for the number of concurrent threads requested
	for i := 0; i < concurrency; i++ {
		// fmt.Printf("Processing messages on thread %v...\n", i)
		go func() {
			for msg := range msgs {
				// if tha handler returns true then ACK, else NACK
				// the message back into the rabbit queue for
				// another round of processing
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
