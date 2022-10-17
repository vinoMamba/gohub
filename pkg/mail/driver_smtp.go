package mail

import (
	"fmt"
	"net/smtp"

	emailPKG "github.com/jordan-wright/email"
	"github.com/vinoMamba/gohub/pkg/logger"
)

type SMTP struct{}

func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()
	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML
	err := e.Send(
		"smtp.gmail.com:587",
		smtp.PlainAuth(
			"",
			config["username"],
			config["password"],
			config["host"],
		),
	)
	if err != nil {
		return false
	}
	logger.Info("Email sent successfully")
	return true
}
