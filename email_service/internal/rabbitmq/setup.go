package rabbitmq

func (c *RabbitMQClient) SetupQueues() error {
	_, err := c.Channel.QueueDeclare(
		"email_queue",
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	return err
}