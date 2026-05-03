package email

type Email struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Html    string `json:"html"`
	Text    string `json:"text"`
}