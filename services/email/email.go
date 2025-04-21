package email

import "os"

// IEmailService is used by other services to send email
type IEmailService interface {
	SendEmail(data ISendEmail) error
}

type emailService struct {
	from       string
	password   string
	smtpServer string
	smtpPort   string
}

// CreateEmailService is a factory to create email service
func CreateEmailService() IEmailService {
	return &emailService{
		from:       os.Getenv("EMAIL"),
		password:   os.Getenv("PASSWORD"),
		smtpServer: os.Getenv("SMTP_SERVER"),
		smtpPort:   os.Getenv("SMTP_PORT"),
	}
}
