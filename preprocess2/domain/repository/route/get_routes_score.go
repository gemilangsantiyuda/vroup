package route

import (
	"math"

	"github.com/vroup/preprocess2/domain/model"
)

//GetRoutesScore get the score for each Objective Function from formed routes
func GetRoutesScore(routeList []model.Route, distanceMap map[string]float64, ratingMap map[string]float64) model.Score {

	//calculating distance
	distance := getTotalDistance(routeList, distanceMap)
	kitchenOptimality := getKitchenOptimality(routeList)
	ratingScore := getRatingScore(routeList, ratingMap)

	score := model.Score{
		TotalDistance:     distance,
		KitchenOptimality: kitchenOptimality,
		RatingScore:       ratingScore,
	}

	return score
}

func getTotalDistance(routeList []model.Route, distanceMap map[string]float64) float64 {
	distance := 0.0
	for _, route := range routeList {
		dist := GetRouteDistance(route, distanceMap)
		distance += dist
	}
	return distance
}

func getKitchenOptimality(routeList []model.Route) float64 {

	kitchenServingQtyMap := make(map[string]int)
	var kitchenList []model.Node
	optimality := 0.0

	for _, route := range routeList {

		kitchen := route[0]
		servingQty := 0

		//check if kitchen is not already indexed in map yet
		//if already then register its val to servingQty
		if _, ok := kitchenServingQtyMap[kitchen.ID]; !ok {
			kitchenServingQtyMap[kitchen.ID] = 0
			kitchenList = append(kitchenList, kitchen)
		} else {
			servingQty = kitchenServingQtyMap[kitchen.ID]
		}

		for idx := 1; idx < len(route); idx++ {
			servingQty += route[idx].Qty
		}
		kitchenServingQtyMap[kitchen.ID] = servingQty
	}

	//Calculate optimality for each kitchen
	for _, kitchen := range kitchenList {

		servingQty := kitchenServingQtyMap[kitchen.ID]
		kitchenCap := kitchen.Capacity
		kitchenOptimality := calculateKitchenOptimality(servingQty, kitchenCap)
		optimality += kitchenOptimality
	}

	optimality = float64(len(kitchenList)) / optimality
	return optimality
}

//this needs to be thought through again!
func calculateKitchenOptimality(servingQty int, kitchenCap model.KitchenCapacity) float64 {

	optimality := math.Abs(float64(servingQty - kitchenCap.Optimum))
	return optimality
}

func getRatingScore(routeList []model.Route, ratingMap map[string]float64) float64 {

	ratingScore := 0.0
	nodeCount := 0

	for _, route := range routeList {

		kitchen := route[0]
		for idx := 1; idx < len(route); idx++ {

			key := route[idx].ID + "-" + kitchen.ID
			rate := ratingMap[key]
			ratingScore += rate
			nodeCount++
		}
	}

	ratingScore /= float64(nodeCount)
	return ratingScore
}
