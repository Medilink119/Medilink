package mail

import (
	"net/smtp"
)

type Mailer struct {
	sender string
	pass   string
	host   string
	auth   smtp.Auth
}

func InitMailer(sender, pass, host string) *Mailer {
	return &Mailer{
		sender: sender,
		pass:   pass,
		host:   host,
		auth:   smtp.PlainAuth("", sender, pass, host),
	}
}

func (m *Mailer) Send(to string) error {
	message := "From: " + m.sender + "\n" + "To: " + to + "\n" + "Subject: Appointment Reminder\n\n"
	var err error
	for i := 0; i < 3; i++ {
		err = smtp.SendMail("smtp.gmail.com:587", m.auth, m.sender, []string{to}, []byte(message))
		if err == nil {
			break
		}
	}
	return err
}
