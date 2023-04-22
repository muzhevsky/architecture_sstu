package utilTypes

type OrderDto struct {
	SellerId   int `json:"SellerId"`
	CustomerId int `json:"CustomerId"`
	OrderId    int `json:"OrderId"`
	Status     int `json:"Status"`
}
