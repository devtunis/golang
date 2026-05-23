package main

import (
	"fmt"
)

// go install github.com/air-verse/air@latest

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// 1. 3f9a8c2d-1b7e-4a6f-9d12-8e5c0a7b9c41
// 2. a7d3f2b1-6c9e-4f0a-b2d8-1e3f7c9a5b62
// 3. 91c4e8fa-2d6b-4b7c-8a1f-5e9d3c2a7f04
// 4. d2a9c6f1-3e7b-4d8a-9f21-6c0b5e7a3d98
// 5. 6b1e4c9f-8a2d-4f7b-b3c1-0d9e6a5f2c73
// 6. f8c3a1d7-5e9b-4c2f-8d6a-3b7e0c1a9f25
// 7. 0d7a5c3e-9f1b-4a8d-b2c6-7e3f1a9c4b80
// 8. c4b2e9a1-7d3f-4c6a-8e1b-9f0d2a5c7e63
// 9. 5e8a1c7d-2b9f-4d3a-a6c1-8f0e7b2c4d91
// 10. b9d6f3a2-1c8e-4a7b-9d5f-2e0c6a3b7f84

func main() {

	var highScore []int

	Rooms := make(map[string][]int)
	Rooms["room1"] = highScore
	Rooms["room1"] = append(Rooms["room1"], 10, 10, 20, 30)

	RoomsStruct := make(map[string][]User)
	RoomsStruct["a"] = append(RoomsStruct["uniqueid"], User{"Gaith", "NahdiGhaith@gmail.com", true, 21})
	RoomsStruct["a"] = append(RoomsStruct["uniqueid"], User{"wajib", "wajih@gmail.com", true, 21})
	RoomsStruct["a"] = append(RoomsStruct["uniqueid"], User{"mourad", "mouradayedi@gmail.com", true, 21})
	RoomsStruct["b"] = append(RoomsStruct["uniqueid"], User{"mourad", "mouradayedi@gmail.com", true, 21})

	num := 10

	if num := 20; num > 30 {
		fmt.Println("Numm bigger than 20")
	} else {
		fmt.Println("not bigger")
	}
	fmt.Println(RoomsStruct)
	fmt.Println(num)

}
