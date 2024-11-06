package services

type Reservation struct {
	Name    string `json:"name"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Guests  int    `json:"guests,string"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message,omitempty"`
}

func (c Reservation) GetEmail() string {
	return c.Email
}

func (c *Reservation) HandleReservationData() error {
	err := sendMail[Reservation](*c)
	if err != nil {
		return err
	}

	return nil
}
