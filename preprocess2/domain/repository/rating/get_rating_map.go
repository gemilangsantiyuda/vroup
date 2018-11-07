package rating

import (
	"strconv"

	"github.com/vroup/preprocess2/domain/model"
)

//GetRatingMap get the rating map with the key of userID-kitchenID
func GetRatingMap(ratingList []model.Rating) map[string]float64 {

	ratingMap := make(map[string]float64)

	for _, rating := range ratingList {

		userID := strconv.Itoa(rating.UserID)
		kitchenID := strconv.Itoa(rating.KitchenID)
		key := userID + "-" + kitchenID
		ratingMap[key] = rating.Rate
	}
	return ratingMap
}
