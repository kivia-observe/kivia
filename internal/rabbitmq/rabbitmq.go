package rabbitmq

import (
	"log"
	logger "log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Conn *amqp.Connection

	Channel *amqp.Channel
}

func NewRabbitMQClient(rabbitMQConnectionUrl string) *RabbitMQClient {

	// CONNECT TO RABBITMQ
	conn, err := amqp.Dial(rabbitMQConnectionUrl)
	if err != nil {
		logger.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	// OPEN A RABBITMQ CHANNEL
	ch, err := conn.Channel()
	if err != nil {
		logger.Fatalf("Failed to open a RabbitMQ channel: %s", err)
	}

	// STORE RABBITMQ CONNECTION AND CHANNEL
	return &RabbitMQClient{
		Conn:    conn,
		Channel: ch,
	}
}

func (r *RabbitMQClient) ConsumeRabbitMQQueue(queue_name string) (<-chan amqp.Delivery, error) {

	q, err := r.Channel.QueueDeclare(
		queue_name, // name of the queue
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)

	if err != nil {
		log.Printf("Failed to declare a RabbitMQ queue: %s", err)
		return nil, err
	}

	// SUBSCRIBE TO THE QUEUE
	msgs, err := r.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		log.Printf("Failed to register a RabbitMQ consumer: %s", err)
		return nil, err
	}

	return msgs, nil
}