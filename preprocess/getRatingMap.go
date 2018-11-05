package main

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

func getRatingMap(ratingCSVPath string, okMap map[string]string) map[string]float64 {

	ratingMap := make(map[string]float64)
	ratingCount := make(map[string]int)
	ratingCSV := readFile(ratingCSVPath)
	r := csv.NewReader(strings.NewReader(ratingCSV))
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		orderID := record[1]
		userID := record[3]
		rating, err := strconv.Atoi(record[4])
		kitchenID := okMap[orderID]

		userXkitchen := userID + "X" + kitchenID

		if _, ok := ratingMap[userXkitchen]; !ok {
			ratingMap[userXkitchen] = float64(rating)
			ratingCount[userXkitchen] = 1
		} else {
			ratingMap[userXkitchen] += float64(rating)
			ratingCount[userXkitchen]++
		}

	}

	for key, _ := range ratingMap {
		ratingMap[key] = ratingMap[key] / float64(ratingCount[key])
	}

	return ratingMap

}
