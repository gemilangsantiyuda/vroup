package main

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"

	"github.com/skripsi/preprocess/model"
)

func getUserMap(userCSVPath string) map[string]model.User {

	userMap := make(map[string]model.User)
	userCSV := readFile(userCSVPath)

	r := csv.NewReader(strings.NewReader(userCSV))
	for i := 1; ; i++ {

		record, err := r.Read()
		if i == 1 {
			continue
		}
		if err == io.EOF {
			break
		}
		checkErr(err)

		latitude, err := strconv.ParseFloat(record[2], 64)
		checkErr(err)
		longitude, err := strconv.ParseFloat(record[3], 64)
		checkErr(err)

		user := model.User{
			record[0],
			latitude,
			longitude,
		}
		userMap[record[0]] = user
	}

	return userMap
}
