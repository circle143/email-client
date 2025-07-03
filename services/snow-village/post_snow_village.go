package snowvillage

import "net/http"

type hNewMessage struct {
	Name    string `validate:"required,min=3"`
	Phone   string `validate:"required,e164"`
	Email   string `validate:"required,email"`
	Message string
}

func (s *snowVillageService) newMessage(w http.ResponseWriter, r *http.Request) {

}
