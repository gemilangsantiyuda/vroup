package rating

import (
	"github.com/vroup/preprocess2/domain/model"
)

//GetRating get the rating of customer toward the kitchen from the ratingMap
func GetRating(orderNode, kitchenNode model.Node, ratingMap map[string]float64) float64 {

	key := orderNode.ID + "-" + kitchenNode.ID
	return ratingMap[key]
}
