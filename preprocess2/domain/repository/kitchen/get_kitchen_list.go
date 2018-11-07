package kitchen

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/vroup/preprocess2/domain/model"
)

//GetKitchenList Get the kitchen list from the csv input
func GetKitchenList(deliveryDate string) ([]model.Kitchen, error) {

	var kitchenList []model.Kitchen

	csvPath := "Kitchen " + deliveryDate + ".csv"
	file, err := ioutil.ReadFile("prepared data/" + csvPath)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(strings.NewReader(string(file)))

	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		latitude, err := strconv.ParseFloat(record[1], 64)
		longitude, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}

		coordinate := model.Coordinate{
			Latitude:  latitude,
			Longitude: longitude,
		}

		minCap, err := strconv.Atoi(record[3])
		optCap, err := strconv.Atoi(record[4])
		maxCap, err := strconv.Atoi(record[5])
		if err != nil {
			return nil, err
		}

		capacity := model.KitchenCapacity{
			Minimum: minCap,
			Optimum: optCap,
			Maximum: maxCap,
		}
		kitchen := model.Kitchen{
			ID:         id,
			Coordinate: coordinate,
			Capacity:   capacity,
		}

		kitchenList = append(kitchenList, kitchen)

	}

	return kitchenList, nil
}
