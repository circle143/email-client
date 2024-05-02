package helper

import (
	"encoding/json"
	"net/http"

	"mennr.tech/api/services"
)

func DecodeJson(w http.ResponseWriter, r *http.Request, placeholder interface{}) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(placeholder)

	if err != nil {
		return err
	}

	return nil
}

func EncodeJson(w http.ResponseWriter, status int, data interface{}) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)

	if err != nil {

		return err
	}
	return nil
}

func ErrorResponse(w http.ResponseWriter, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload services.JSONResponse

	payload.Error = true
	payload.Message = err.Error()

	EncodeJson(w, statusCode, payload)
}
