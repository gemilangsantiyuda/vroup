package main

import "fmt"

func main1() {

	path := "/home/gemilang/go/src/github.com/skripsi/preprocess/data/"
	chosenDate := "2018-01-19"

	userCSVPath := path + "users.csv"
	omfCSVPath := path + "order_map_forecast.csv"
	ouCSVPath := path + "order_user.csv"
	kitchenCSVPath := path + "kitchen.csv"
	kitchenCapCSVPath := path + "kitchen_capacities.csv"
	ratingCSVPath := path + "rating.csv"

	userMap := getUserMap(userCSVPath)
	ouMap := getOuMap(ouCSVPath)
	okMap := getOkMap(omfCSVPath)
	orderList := getOrderList(omfCSVPath, ouMap, userMap, chosenDate)
	kitchenMap := getKitchenMap(kitchenCSVPath)
	kitchenList := getKitchenList(kitchenCapCSVPath, kitchenMap, chosenDate)

	ratingMap := getRatingMap(ratingCSVPath, okMap)
	chosenRatingList := getChosenRatingList(ratingMap, orderList, kitchenList)

	fmt.Println(kitchenList)
	createOrderFile(orderList, chosenDate)
	createKitchenFile(kitchenList, chosenDate)
	createRatingFile(chosenRatingList, chosenDate)
	fmt.Print(chosenRatingList)
	fmt.Print("FINISH")

}
