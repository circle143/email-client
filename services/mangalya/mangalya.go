package mangalya

import (
	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/common"
)

type mangalyaService struct {
	emailService email.IEmailService
}

func CreateMangalyaService(app common.IApp) common.IService {
	return &mangalyaService{
		emailService: app.GetEmailService(),
	}
}
