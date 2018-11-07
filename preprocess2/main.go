package main

import (
	"fmt"
	"log"

	"github.com/vroup/preprocess2/domain/model"
	kitchenRepo "github.com/vroup/preprocess2/domain/repository/kitchen"
	nodeRepo "github.com/vroup/preprocess2/domain/repository/node"
	orderRepo "github.com/vroup/preprocess2/domain/repository/order"
	ratingRepo "github.com/vroup/preprocess2/domain/repository/rating"
)

func main() {

	chosenDate := "2017-08-16"
	nodeList, ratingMap, distanceMap, err := prepareInput(chosenDate)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(nodeList)
	fmt.Print(ratingMap)
	fmt.Print(distanceMap)

}

func prepareInput(chosenDate string) ([]model.Node, map[string]float64, map[string]float64, error) {

	kitchenList, err := kitchenRepo.GetKitchenList(chosenDate)
	if err != nil {
		return nil, nil, nil, err
	}
	fmt.Println(kitchenList)

	orderList, err := orderRepo.GetOrderList(chosenDate)
	if err != nil {
		return nil, nil, nil, err
	}

	ratingList, err := ratingRepo.GetRatingList(chosenDate)
	if err != nil {
		return nil, nil, nil, err
	}
	ratingMap := ratingRepo.GetRatingMap(ratingList)

	//These 3 are the real input to the algorithm
	orderRatingMap := ratingRepo.GetOrderRatingMap(ratingMap, orderList, kitchenList)
	nodeList := nodeRepo.GetNodeList(orderList, kitchenList)
	distanceMap := make(map[string]float64)

	return nodeList, distanceMap, orderRatingMap, nil
}
