package model

type Order struct {
	OrderID             string
	UserID              string
	Qty                 int
	Latitude, Longitude float64
}
