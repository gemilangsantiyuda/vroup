package order

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/vroup/preprocess2/domain/model"
)

//GetOrderList get the list of orders from the csv input
func GetOrderList(deliveryDate string) ([]model.Order, error) {

	var orderList []model.Order

	csvPath := "Order " + deliveryDate + ".csv"
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

		userID, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}

		qty, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, err
		}

		latitude, err := strconv.ParseFloat(record[3], 64)
		longitude, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return nil, err
		}

		coordinate := model.Coordinate{
			Latitude:  latitude,
			Longitude: longitude,
		}

		order := model.Order{
			ID:         id,
			UserID:     userID,
			Qty:        qty,
			Coordinate: coordinate,
		}

		orderList = append(orderList, order)

	}

	return orderList, nil
}
