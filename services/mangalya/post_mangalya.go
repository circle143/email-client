package mangalya

import (
	"circledigital.in/api/services/email"
	"circledigital.in/api/utils/custom"
	"circledigital.in/api/utils/payload"
	"net/http"
)

type hNewFormSubmission struct {
	Name     string `validate:"required,min=3"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,e164"`
	Location string `validate:"required"`
	More     string
}

func (nfs *hNewFormSubmission) GetToSend() []string {
	return []string{"mangalyagroup@gmail.com", nfs.Email}
}

func (nfs *hNewFormSubmission) GetSubject() string {
	return "Mangalya Enquiry Details"
}

func (nfs *hNewFormSubmission) GetTemplateDir() string {
	return "./assets/mangalya/template.html"
}

func (nfs *hNewFormSubmission) GetTemplateData() any {
	// logo and hero should be the same as attachment contentId
	return struct {
		Data *hNewFormSubmission
		Logo string
		Hero string
	}{
		Data: nfs,
		Logo: "logo",
	}
}

func (nfs *hNewFormSubmission) GetAttachments() []email.Attachment {
	return []email.Attachment{
		{
			Path:        "./assets/mangalya/logo.png",
			ContentType: "image/png",
			ContentID:   "logo",
			Inline:      true,
		},
	}
}

func (ms *mangalyaService) newFormSubmission(w http.ResponseWriter, r *http.Request) {
	details := payload.ValidateAndDecodeRequest[hNewFormSubmission](w, r)

	err := ms.emailService.SendEmail(details)
	if err != nil {
		payload.HandleError(w, err)
		return
	}

	var response custom.JSONResponse
	response.Error = false
	response.Message = "Successfully received your details."

	payload.EncodeJSON(w, http.StatusAccepted, response)
}