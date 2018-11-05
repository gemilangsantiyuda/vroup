package main

import (
	"encoding/csv"
	"io"
	"strings"
)

func getOkMap(omfCSVPath string) map[string]string {

	okMap := make(map[string]string)

	omfCSV := readFile(omfCSVPath)
	r := csv.NewReader(strings.NewReader(omfCSV))
	for {

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			continue
		}

		okMap[record[1]] = record[2]
	}
	return okMap
}
