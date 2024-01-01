package email

type Emailer interface {
	sendEmail(subject string, recepients []string, body []byte, cc []string, bcc []string) error
	ValidateEmail(recepient string) (string, error)
	ChangePasswordConfirm(recepient string) error
}

type Email struct {
	SenderEmail    string
	SenderPassword string
	SmtpAddress    string
	SmtpPort       string
}

func NewEmail(senderEmail, senderPassword string) Emailer {
	return &Email{
		SenderEmail:    senderEmail,
		SenderPassword: senderPassword,
	}
}
