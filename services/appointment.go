package services

type Appointment struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	VehicleYear        int    `json:"vehicleYear,string"`
	VehicleMake        string `json:"vehicleMake"`
	VehicleModel       string `json:"vehicleModel"`
	VehicleRequirement string `json:"vehicleRequirement,omitempty"`
	RequestedDate      string `json:"requestedDate"`
	RequestedTime      string `json:"requestedTime"`
}

func (a Appointment) GetEmail() string {
	return a.Email
}

func (a *Appointment) HandleAppointmentData() error {
	err := sendMail[Appointment](*a)

	if err != nil {
		return err
	}

	return nil
}
