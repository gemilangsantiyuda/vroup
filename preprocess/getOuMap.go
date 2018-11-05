package main

import (
  "encoding/csv"
  "strings"
  "io"
)

func getOuMap(ouCSVPath string) map[string]string {

  ouMap := make(map[string]string)

  ouCSV := readFile(ouCSVPath)
  r := csv.NewReader(strings.NewReader(ouCSV))

  for {

    record,err := r.Read()
    if err == io.EOF{
      break
    }
    checkErr(err)

    ouMap[record[0]]=record[1]
  }

  return ouMap
}
