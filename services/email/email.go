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

type AuthenticationDetails struct {
	Email    string
	Password string
}

// CreateEmailServiceWithCustomAuth creates email service with custom auth details
func CreateEmailServiceWithCustomAuth(auth AuthenticationDetails) IEmailService {
	return &emailService{
		from:       auth.Email,
		password:   auth.Password,
		smtpServer: os.Getenv("SMTP_SERVER"),
		smtpPort:   os.Getenv("SMTP_PORT"),
	}
}
