package main

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/skripsi/preprocess/model"
)

func createRatingFile(chosenRatingList []model.Rating, chosenDate string) {

	file, err := os.Create("Rating " + chosenDate + ".csv")
	checkErr(err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, rating := range chosenRatingList {

		userID := rating.UserID
		kitchenID := rating.KitchenID
		rate := strconv.FormatFloat(rating.Rating,'f',-1,64)
		err := writer.Write([]string{userID, kitchenID, rate})
		checkErr(err)
	}
}
