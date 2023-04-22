package utilTypes

type status int8

const (
	Pending          status = 0
	SellerRejected   status = -1
	PaymentRejected  status = -2
	SellerSucceeded  status = 1
	PaymentSucceeded status = 2
)
