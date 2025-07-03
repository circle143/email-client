package snowvillage

import (
	"net/http"

	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/custom"
	"circledigital.in/api/utils/payload"
)

type hNewMessage struct {
	Name    string `validate:"required,min=3"`
	Phone   string `validate:"required,e164"`
	Email   string `validate:"required,email"`
	Origin  string `validate:"required"`
	Message string
}

func (h *hNewMessage) validate() bool {
	userOriginValue := userOrigin(h.Origin)
	return userOriginValue.IsValid()
}

func (h *hNewMessage) GetToSend() []string {
	return []string{"snowvillage@cozylounge.in", h.Email}
}

func (h *hNewMessage) GetSubject() string {
	return "Snow Village New Message"
}

func (h *hNewMessage) GetTemplateDir() string {
	return "./assets/snow-village/template.html"
}

func (h *hNewMessage) GetTemplateData() any {
	return struct {
		Data *hNewMessage
		Logo string
		Hero string
	}{
		Data: h,
		Logo: "logo",
		Hero: "hero",
	}
}

func (h *hNewMessage) GetAttachments() []email.Attachment {
	return []email.Attachment{
		{
			Path:        "./assets/snow-village/logo.png",
			ContentType: "image/png",
			ContentID:   "logo",
			Inline:      true,
		},
		{
			Path:        "./assets/snow-village/hero.png",
			ContentType: "image/png",
			ContentID:   "hero",
			Inline:      true,
		},
	}
}

func (s *snowVillageService) newMessage(w http.ResponseWriter, r *http.Request) {
	message := payload.ValidateAndDecodeRequest[hNewMessage](w, r)
	if message == nil {
		return
	}

	if !message.validate() {
		userOriginError := custom.RequestError{
			Status:  http.StatusBadRequest,
			Message: "Invalid value for user-orign",
		}
		payload.HandleError(w, userOriginError)
		return
	}

	err := s.emailService.SendEmail(message)
	if err != nil {
		payload.HandleError(w, err)
		return
	}

	var response custom.JSONResponse
	response.Error = false
	response.Message = "Successfully received your message."

	payload.EncodeJSON(w, http.StatusAccepted, response)
}
