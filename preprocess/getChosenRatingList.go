package main

import (
  "github.com/skripsi/preprocess/model"
)


func getChosenRatingList(ratingMap map[string]float64, orderList []model.Order, kitchenList []model.Kitchen) []model.Rating{

  var chosenRatingList []model.Rating
  for _, order := range orderList{
    for _, kitchen := range kitchenList{

      rating := model.Rating{
        order.UserID,
        kitchen.KitchenID,
        3,
      }

      userXkitchen := order.UserID + "X" + kitchen.KitchenID
      if val, ok := ratingMap[userXkitchen]; ok{
          rating.Rating = val
      }

      chosenRatingList = append(chosenRatingList,rating)
    }
  }

  return chosenRatingList
}
