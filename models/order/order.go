package order

type Order struct {
	OrderId    uint32  `json:"order_id"`
	UserId     uint32  `json:"user_id"`
	DriverId   uint32  `json:"driver_id"`
	From       string  `json:"from"`
	To         string  `json:"to"`
	TaxiType   string  `json:"taxi_type"`
	Date       string  `json:"date"`
	Status     string  `json:"status"`
	UserRate   float32 `json:"user_rate"`
	DriverRate float32 `json:"driver_rate"`
}

type UserRequest struct {
	ID       int32  `json:"id"`
	TaxiType string `json:"taxi_type"`
	From     string `json:"from"`
	To       string `json:"to"`
}
