package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/skripsi/preprocess/model"
)

func createKitchenFile(kitchenList []model.Kitchen, chosenDate string) {

	file, err := os.Create("Kitchen " + chosenDate + ".csv")
	checkErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, kitchen := range kitchenList {

		kitchenID := kitchen.KitchenID
		latitude := strconv.FormatFloat(kitchen.Latitude, 'f', -1, 64)
		longitude := strconv.FormatFloat(kitchen.Longitude, 'f', -1, 64)
		minCap := strconv.Itoa(kitchen.MinCapacity)
		optCap := strconv.Itoa(kitchen.OptCapacity)
		maxCap := strconv.Itoa(kitchen.MaxCapacity)

		err := writer.Write([]string{kitchenID, latitude, longitude, minCap, optCap, maxCap})
		checkErr(err)
	}
}
