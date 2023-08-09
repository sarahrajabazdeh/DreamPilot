package email

import (
	"net/smtp"

	"github.com/sarahrajabazdeh/DreamPilot/config"
)

func SendEmail(to, subject, body string) error {
	emailConfig := config.Config.EmailConfig // Access email configuration from your config package

	auth := smtp.PlainAuth("", emailConfig.SMTPUsername, emailConfig.SMTPPassword, emailConfig.SMTPServer)

	msg := "To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=UTF-8" + "\r\n" +
		"\r\n" + body

	return smtp.SendMail(emailConfig.SMTPServer+":"+string(emailConfig.SMTPPort), auth, emailConfig.SenderEmail, []string{to}, []byte(msg))
}
