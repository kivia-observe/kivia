package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQClient(url string) *RabbitMQClient {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}

	return &RabbitMQClient{
		Connection: conn,
		Channel:    ch,
	}
}

func (r *RabbitMQClient) Consume(queue_name string) (<- chan amqp.Delivery, error){
	
	msgs, err := r.Channel.Consume(queue_name, "", true, false, false, false, nil)
	
	if err != nil {
		
		log.Fatalf("Failed to consume from queue %s: %v", queue_name, err)
		return nil, err
	}
	
	return msgs, nil
} 

func (r *RabbitMQClient) Close(){

	if err := r.Channel.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ channel: %s", err)
	}

	if err := r.Connection.Close(); err != nil {
		log.Printf("Failed to close RabbitMQ connection: %s", err)
	}
}