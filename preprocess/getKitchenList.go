package main
import(
  "github.com/skripsi/preprocess/model"
  "encoding/csv"
  "strings"
  "io"
  "strconv"
)
func getKitchenList(kitchenCapCSVPath string, kitchenMap map[string]model.Kitchen,chosenDate string) []model.Kitchen{

  var kitchenList []model.Kitchen

  kitchenCapCSV := readFile(kitchenCapCSVPath)
  r := csv.NewReader(strings.NewReader(kitchenCapCSV))

  for {
    record, err := r.Read()
    if err == io.EOF{
      break
    }
    if err!=nil{
      continue
    }
    if record[3]!= "1" || record[1]!=chosenDate{
      continue
    }

    kitchenID := record[2]
    kitchen := kitchenMap[kitchenID]
    kitchen.MinCapacity,err = strconv.Atoi(record[6])
    kitchen.OptCapacity,err = strconv.Atoi(record[5])
    kitchen.MaxCapacity,err = strconv.Atoi(record[4])

    kitchenList = append(kitchenList, kitchen)
  }

  return kitchenList
}
