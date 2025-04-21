package cozy

import (
	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/common"
	"os"
)

type cozyService struct {
	emailService email.IEmailService
}

func CreateCozyService(_ common.IApp) common.IService {
	return &cozyService{
		emailService: email.CreateEmailServiceWithCustomAuth(email.AuthenticationDetails{
			Email:    os.Getenv("COZY_EMAIL"),
			Password: os.Getenv("COZY_PASSWORD"),
		}),
	}
}
