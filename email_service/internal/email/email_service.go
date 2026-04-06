package email

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/winnerx0/kivia/email_service/internal/rabbitmq"
)

type emailservice struct {
	brevoApiKey    string
	rabbitMQClient *rabbitmq.RabbitMQClient
	senderEmail    string
	senderName     string
}

func NewEmailService(brevoApiKey string, senderEmail string, senderName string, rabbitMQClient *rabbitmq.RabbitMQClient) *emailservice {
	return &emailservice{brevoApiKey: brevoApiKey, senderEmail: senderEmail, senderName: senderName, rabbitMQClient: rabbitMQClient}
}

func (s emailservice) SendEmail(email Email) error {

	emailBytes, err := json.Marshal(email)

	if err != nil {
		return err
	}

	err = s.rabbitMQClient.Channel.Publish("", "email_queue", false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         emailBytes,
		DeliveryMode: amqp.Persistent,
	})

	if err != nil {
		return err
	}

	return nil
}

func (s emailservice) EmailConsumer(ctx context.Context) error {

	log.Println("sending the email")

	messages, err := s.rabbitMQClient.Consume("email_queue")

	if err != nil {
		return err
	}

	go func() {

		for {
			select {
			case <-ctx.Done():
				log.Printf("Context cancelled, stopping email service")
				return
			case msg, ok := <-messages:
				if !ok {
					log.Printf("Channel closed")
					return
				}

				var email Email
				err = json.Unmarshal(msg.Body, &email)

				if err != nil {
					log.Printf("Failed to unmarshal email: %v", err)
					continue
				}

				data, err := json.Marshal(map[string]any{
					"sender": struct {
						Email string `json:"email"`
						Name  string `json:"name"`
					}{
						Email: s.senderEmail,
						Name:  s.senderName,
					},
					"to": []struct {
						Email string `json:"email"`
					}{
						{Email: email.To},
					},
					"subject":     email.Subject,
					"htmlContent": email.Html,
					"textContent": email.Text,
				})

				if err != nil {
					log.Printf("Failed to marshal email: %v", err)
					continue
				}

				req, err := http.NewRequest(http.MethodPost, "https://api.brevo.com/v3/smtp/email", bytes.NewReader(data))

				if err != nil {
					log.Printf("Failed to create request: %v", err)
					continue
				}

				req.Header.Set("api-key", s.brevoApiKey)
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)

				if err != nil {
					log.Printf("Failed to send email: %v", err)
					continue
				}

				if resp.StatusCode != http.StatusCreated {
					log.Printf("Failed to send email: %s", resp.Status)
					continue
				}

				log.Printf("Sending email to: %s, subject: %s, body: %s", email.To, email.Subject, email.Text)

			}
		}
	}()

	return nil
}
