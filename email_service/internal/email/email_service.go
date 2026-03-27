package email

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/winnerx0/dyno/email_service/internal/rabbitmq"
)

type emailservice struct {
	brevoApiKey    string
	rabbitMQClient *rabbitmq.RabbitMQClient
}

func NewEmailService(brevoApiKey string, rabbitMQClient *rabbitmq.RabbitMQClient) *emailservice {
	return &emailservice{brevoApiKey: brevoApiKey, rabbitMQClient: rabbitMQClient}
}

func (s emailservice) SendEmail(email Email) error {
	
	emailBytes, err := json.Marshal(email)
	
	if err != nil {
		return err
	}

	err = s.rabbitMQClient.Channel.Publish("", "email_queue", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body: emailBytes,
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

				reqBody := map[string]string{
					"sender":      "eragbonikpeya@gmail.com",
					"name":        "Dyno",
					"subject":     email.Subject,
					"htmlContent": email.Body,
					"to":          email.To,
				}

				jsonBody, err := json.Marshal(reqBody)

				if err != nil {
					log.Printf("Failed to marshal email: %v", err)
					continue
				}

				req, _ := http.NewRequest(http.MethodPost, "https://api.brevo.com/v3/smtp/email", bytes.NewReader(jsonBody))

				req.Header.Add("api-key", s.brevoApiKey)
				req.Header.Add("Content-Type", "application/json")
				req.Header.Set("Accept", "application/json")

				res, err := http.DefaultClient.Do(req)

				if err != nil {
					log.Printf("Failed to send email: %v", err)
					continue
				}

				defer res.Body.Close()

				bodyBytes, err := io.ReadAll(res.Body)

				if err != nil && err != io.EOF {
					log.Printf("Failed to read email response: %v", err)
					continue
				}
				fmt.Println(res)
				fmt.Println(string(bodyBytes))

				log.Printf("Sending email to: %s, subject: %s, body: %s", email.To, email.Subject, email.Body)

			}
		}
	}()

	return nil
}
