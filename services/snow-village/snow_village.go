package snowvillage

import (
	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/common"
)

type snowVillageService struct {
	emailService email.IEmailService
}

func CreateSnowVillageService(app common.IApp) common.IService {
	return &snowVillageService{
		emailService: app.GetEmailService(),
	}
}
