package controllers

import (
	"errors"
	"log"
	"net/http"

	"mennr.tech/api/helper"
	"mennr.tech/api/services"
)

func PostAppointment(w http.ResponseWriter, r *http.Request) {
	data, err := helper.DecodeJson[services.Appointment](w, r)
	if err != nil {
		log.Println(err)
		helper.ErrorResponse(w, err)
		return
	}

	err = data.HandleAppointmentData()

	if err != nil {
		err = errors.New("500 internal server error")
		helper.ErrorResponse(w, err, http.StatusInternalServerError)
		return

	}
	var response services.JSONResponse
	response.Error = false
	response.Message = "successfully received the details. Details we got."
	response.Data = data
	helper.EncodeJson(w, http.StatusAccepted, response)

}
