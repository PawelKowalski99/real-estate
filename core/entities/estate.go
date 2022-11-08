package entities

type Estate struct {
	URL                 string  `json:"url,omitempty"`
	Address             string  `json:"address,omitempty" json:"address,omitempty"`
	DistanceFromCentrum float64 `json:"distance_from_centrum,omitempty" json:"distance_from_centrum,omitempty"`
	// Surface in m^2
	Surface              string `json:"surface,omitempty" json:"surface,omitempty"`
	RoomAmount           string `json:"room_amount,omitempty" json:"room_amount,omitempty"`
	Floor                int    `json:"floor,omitempty" json:"floor,omitempty"`
	IsOutsideParkingSlot bool   `json:"is_outside_parking_slot,omitempty" json:"is_outside_parking_slot,omitempty"`
	IsInsideParkingSlot  bool   `json:"is_inside_parking_slot,omitempty" json:"is_inside_parking_slot,omitempty"`

	PricePerM2 string `json:"price_per_m_2,omitempty" json:"price_per_m_2,omitempty"`

	// To renovate/
	Status string `json:"status,omitempty" json:"status,omitempty"`

	Price string `json:"price,omitempty" json:"price,omitempty"`

	RentPrice string `json:"rent_price,omitempty"`
}
