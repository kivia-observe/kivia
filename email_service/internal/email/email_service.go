package email

import "log"

type emailservice struct {

}

func NewEmailService() *emailservice {
	return &emailservice{}
}

func (s emailservice) SendEmail(to string, subject string, body string) error {

	log.Printf("Sending email to: %s, subject: %s, body: %s", to, subject, body)
	return nil
}