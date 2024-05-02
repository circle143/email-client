package services

type Appointment struct {
	Name               string `json:"name"`
	Email              string `json:"email"`
	PhoneNumber        string `json:"phoneNumber"`
	VehicleYear        int    `json:"vehicleYear,string"`
	VehicleMake        string `json:"vehicleMake"`
	VehicleModel       string `json:"vehicleModel"`
	VehicleRequirement string `json:"vehicleRequirement,omitempty"`
}
