package rating

import (
	"strconv"

	"github.com/vroup/preprocess2/domain/model"
)

//GetOrderRatingMap Get a rating map where order.ID-kitchen.ID is the key
func GetOrderRatingMap(ratingMap map[string]float64, orderList []model.Order, kitchenList []model.Kitchen) map[string]float64 {

	orderRatingMap := make(map[string]float64)

	for _, order := range orderList {

		orderID := strconv.Itoa(order.ID)
		userID := strconv.Itoa(order.UserID)
		for _, kitchen := range kitchenList {

			kitchenID := strconv.Itoa(kitchen.ID)
			key := orderID + "-" + kitchenID
			userRateKey := userID + "-" + kitchenID
			orderRatingMap[key] = ratingMap[userRateKey]
		}
	}

	return orderRatingMap
}
