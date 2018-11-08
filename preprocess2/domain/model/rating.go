package model

//Rating just a struct to store user rating, only used in GetRatingList
type Rating struct {
	UserID    int
	KitchenID int
	Rate      float64
}
