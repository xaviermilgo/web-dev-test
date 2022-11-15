package processors

type Hotel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Driver struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Trip struct {
	Id           string  `json:"id"`
	DriverId     string  `json:"driver_id"`
	HotelId      string  `json:"hotel_id"`
	HotelRating  float64 `json:"hotel_rating"`
	DriverRating float64 `json:"driver_rating"`
	Status       string  `json:"status"`
	Driver       *Driver `json:"-"`
	Hotel        *Hotel  `json:"-"`
}

type TripsData struct {
	Drivers []*Driver  `json:"drivers"`
	Hotels  []*Hotel   `json:"hotels"`
	Trips   chan *Trip `json:"-"`
}
