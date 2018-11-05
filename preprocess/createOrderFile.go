package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/skripsi/preprocess/model"
)

func createOrderFile(orderList []model.Order, chosenDate string) {

	file, err := os.Create("Order " + chosenDate + ".csv")
	checkErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, order := range orderList {

		orderID := order.OrderID
		userID := order.UserID
		qty := strconv.Itoa(order.Qty)
		latitude := strconv.FormatFloat(order.Latitude, 'f', -1, 64)
		longitude := strconv.FormatFloat(order.Longitude, 'f', -1, 64)
		err := writer.Write([]string{orderID, userID, qty, latitude, longitude})
		checkErr(err)
	}
}
