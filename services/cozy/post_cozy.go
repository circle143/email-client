package cozy

import (
	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/custom"
	"circledigital.in/api/utils/payload"
	"net/http"
)

type hAddNewReservation struct {
	Name    string `validate:"required,min=3"`
	Date    string `validate:"required"`
	Time    string `validate:"required"`
	Guests  int    `validate:"required,gte=1,lte=8"`
	Email   string `validate:"required,email"`
	Phone   string `validate:"required,e164"`
	Message string
}

func (nr *hAddNewReservation) GetToSend() []string {
	return []string{"info@cozylounge.in", nr.Email}
}

func (nr *hAddNewReservation) GetSubject() string {
	return "Cozy lounge reservation"
}

func (nr *hAddNewReservation) GetTemplateDir() string {
	return "./assets/cozy/template.html"
}

func (nr *hAddNewReservation) GetTemplateData() any {
	// logo and hero should be the same as attachment contentId
	return struct {
		Data *hAddNewReservation
		Logo string
		Hero string
	}{
		Data: nr,
		Logo: "logo",
		Hero: "hero",
	}
}

func (nr *hAddNewReservation) GetAttachments() []email.Attachment {
	return []email.Attachment{
		{
			Path:        "./assets/cozy/logo.png",
			ContentType: "image/png",
			ContentID:   "logo",
			Inline:      true,
		},
		{
			Path:        "./assets/cozy/hero.png",
			ContentType: "image/png",
			ContentID:   "hero",
			Inline:      true,
		},
	}
}

func (cs *cozyService) addNewReservation(w http.ResponseWriter, r *http.Request) {
	reservation := payload.ValidateAndDecodeRequest[hAddNewReservation](w, r)
	if reservation == nil {
		return
	}

	err := cs.emailService.SendEmail(reservation)
	if err != nil {
		payload.HandleError(w, err)
		return
	}

	var response custom.JSONResponse
	response.Error = false
	response.Message = "Successfully received reservation details🥳🥳."

	payload.EncodeJSON(w, http.StatusAccepted, response)
}
