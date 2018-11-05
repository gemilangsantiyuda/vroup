package main

import "io/ioutil"

func readFile(path string) string {

	dat, err := ioutil.ReadFile(path)
	checkErr(err)

	return string(dat)
}
