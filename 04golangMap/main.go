package main

import (
	"fmt"
)

func main() {

	Persons := make(map[string]int)
	Persons["V1"] = 1
	Persons["V2"] = 2
	Persons["V3"] = 3
	Persons["V4"] = 4
	Persons["V5"] = 5

	for key, value := range Persons {
		fmt.Println(key, value)
	}

}
