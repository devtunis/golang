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

	// nested map
	Rooms := make(map[string]map[string]int)
	Rooms["room1"] = map[string]int{"players": 5}
	Rooms["room2"] = map[string]int{"players": 15, "palyer": 19}
	Rooms["room3"] = map[string]int{"player2": 20}

	fmt.Println(Rooms)
	fmt.Println(Rooms["room1"])

}
