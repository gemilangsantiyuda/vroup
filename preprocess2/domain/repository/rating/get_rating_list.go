package rating

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/vroup/preprocess2/domain/model"
)

//GetRatingList get the rating list from the csv input
func GetRatingList(deliveryDate string) ([]model.Rating, error) {

	var ratingList []model.Rating

	csvPath := "Rating " + deliveryDate + ".csv"
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

		userID, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		kitchenID, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, err
		}

		rate, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}

		rating := model.Rating{
			UserID:    userID,
			KitchenID: kitchenID,
			Rate:      rate,
		}

		ratingList = append(ratingList, rating)

	}

	return ratingList, nil
}
