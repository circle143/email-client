package services

type Contact struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	VehicleYear        int    `json:"vehicleYear,string"`
	VehicleMake        string `json:"vehicleMake"`
	VehicleModel       string `json:"vehicleModel"`
	VehicleRequirement string `json:"vehicleRequirement,omitempty"`
}

func (c *Contact) HandleContactData() error {
	err := sendMail(*c)
	if err != nil {
		return err
	}

	return nil
}
