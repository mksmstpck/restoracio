package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/mksmstpck/restoracio/internal/config"
	"github.com/pborman/uuid"
)

func (sender *Email) sendEmail(
	subject string,
	recepients []string,
	body []byte,
	cc []string,
	bcc []string,
) error {
	e := email.NewEmail()
	e.From = sender.SenderEmail
	e.Subject = subject
	e.To = recepients
	e.HTML = body
	e.Cc = cc
	e.Bcc = bcc

	auth := smtp.PlainAuth("", sender.SenderEmail, sender.SenderPassword, sender.SmtpAddress)
	return e.Send(sender.SmtpAddress+":"+sender.SmtpPort, auth)
}

func (sender *Email) ValidateEmail(recepient string) (string, error) {
	id := uuid.NewUUID().String()
	link := config.NewConfig().GlobalURL + "/verify/" + id

	template :=
		`
	<h1>You are trying to register on Restoracio.</h1>
	<h3>Verify your email address.</h3>
	<p>Please click the link below to verify your email address.</p>
	%d
	`

	content := fmt.Sprintf(template, link)
	return id, sender.sendEmail("Verify your email address",
		[]string{recepient},
		[]byte(content),
		nil,
		nil,
	)
}

func (e *Email) ChangePasswordConfirm(recepient string) error
