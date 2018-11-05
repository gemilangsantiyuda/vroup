package main

import(
  "github.com/skripsi/preprocess/model"
  "encoding/csv"
  "strings"
  "io"
  "fmt"
  "strconv"
)

func getKitchenMap(kitchenCSVPath string) map[string]model.Kitchen{

  kitchenMap := make(map[string]model.Kitchen)
  kitchenCSV := readFile(kitchenCSVPath)
  r := csv.NewReader(strings.NewReader(kitchenCSV))

  for {

    record, err := r.Read()
    if err == io.EOF{
      break
    }
    if err!=nil{
      continue
    }

    kitchenID := record[0]
    latitude,err := strconv.ParseFloat(record[3], 64)
    longitude,err := strconv.ParseFloat(record[2], 64)

    kitchen := model.Kitchen{
      kitchenID,
      latitude,
      longitude,
      0,
      0,
      0,
    }
    fmt.Println(kitchen)
    kitchenMap[kitchenID]=kitchen
  }
  return kitchenMap
}
