package email

import (
	"context"
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/winnerx0/dyno/email_service/internal/rabbitmq"
	mail "github.com/wneessen/go-mail"
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
				
				message := mail.NewMsg()
				if err := message.From(os.Getenv("SMTP_USER")); err != nil {
					log.Fatalf("failed to set FROM address: %s", err)
				}
				if err := message.To(email.To); err != nil {
					log.Fatalf("failed to set TO address: %s", err)
				}
				message.Subject(email.Subject)
				message.SetBodyString(mail.TypeTextHTML, email.Body)

				// Deliver the mails via SMTP
				client, err := mail.NewClient("smtp.gmail.com",
					mail.WithSMTPAuth(mail.SMTPAuthAutoDiscover), mail.WithTLSPortPolicy(mail.TLSMandatory),
					mail.WithUsername(os.Getenv("SMTP_USER")), mail.WithPassword(os.Getenv("SMTP_PASS")),
				)
				if err != nil {
					log.Fatalf("failed to create new mail delivery client: %s", err)
				}
				if err := client.DialAndSend(message); err != nil {
					log.Fatalf("failed to deliver mail: %s", err)
				}
				log.Printf("Sending email to: %s, subject: %s, body: %s", email.To, email.Subject, email.Body)

			}
		}
	}()

	return nil
}
