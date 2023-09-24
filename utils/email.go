package utils

import (
	"net/smtp"

	"github.com/mksmstpck/restoracio/internal/config"
	"github.com/pborman/uuid"
)

func EmailValidator(resiver string, id uuid.UUID) error {
	config := config.NewConfig()

	auth := smtp.PlainAuth("", config.EmailSender, config.EmailPassword, config.SMTPHost)
	message := []byte(config.GlobalURL + "/admin/validate/" + id.String())

	err := smtp.SendMail(
		config.SMTPHost+":"+config.SMTPPort,
		auth, config.EmailSender,
		[]string{resiver},
		message,
	)
	if err != nil {
		return err
	}
	return nil
}