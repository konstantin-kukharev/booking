package entity

type Order struct {
	ID        string `json:"id,omitempty"`
	HotelID   string `json:"hotel_id"`
	RoomID    string `json:"room_id"`
	UserEmail string `json:"email"`
	DateInterval
}

func (o *Order) Validate() error {
	if err := o.DateInterval.Validate(); err != nil {
		return err
	}

	return nil
}
