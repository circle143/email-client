package init

import "circledigital.in/api/services/email"

func (a *app) createEmailService() email.IEmailService {
	return email.CreateEmailService()
}
