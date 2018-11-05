package main

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/skripsi/preprocess/model"
)

func getOrderList(omfCSVPath string, ouMap map[string]string, userMap map[string]model.User, chosenDate string) []model.Order {

	var orderList []model.Order
	omfCSV := readFile(omfCSVPath)

	r := csv.NewReader(strings.NewReader(omfCSV))

	for i := 1; ; i++ {

		record, err := r.Read()
		if i == 1 {
			continue
		} else if err == io.EOF {
			break
		} else if err != nil {
			continue
		}
		if record[4] != chosenDate {
			continue
		}

		orderID := record[1]
		userID := ouMap[orderID]
		qty, err := strconv.Atoi(record[3])
		checkErr(err)

		if _, ok := userMap[userID]; !ok {
			continue
		}

		latitude := userMap[userID].Latitude
		longitude := userMap[userID].Longitude

		order := model.Order{
			orderID,
			userID,
			qty,
			latitude,
			longitude,
		}
		orderList = append(orderList, order)
	}
	return orderList
}
