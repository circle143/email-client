package init

import (
	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/common"
	"github.com/go-chi/chi/v5"
)

type app struct {
	mux          *chi.Mux
	emailService email.IEmailService
}

func (a *app) GetRouter() *chi.Mux {
	return a.mux
}

func (a *app) GetEmailService() email.IEmailService {
	return a.emailService
}

func (a *app) initApplication() {
	a.emailService = a.createEmailService()
	a.mux = a.createRouter()
}

func GetApplication() common.IApp {
	appObj := &app{}
	appObj.initApplication()

	return appObj
}
