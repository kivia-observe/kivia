package rabbitmq

func (c *RabbitMQClient) SetupQueues() error {
	_, err := c.Channel.QueueDeclare(
		"log_queue",
		true,  // durable
		false, // auto-delete
		false, // exclusive
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	_, err = c.Channel.QueueDeclare(
		"email_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	return err
}