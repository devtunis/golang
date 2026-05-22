package main

import (
	"fmt"
	"sort"
)

func increamntByPointer(ptr *int) {
	*ptr = *ptr + 200
}
func main() {
	// pointers
	p := fmt.Println

	var ptr *int

	myNumber := 23

	ptr = &myNumber
	increamntByPointer(ptr)
	p(*ptr, "<== pointer")
	//arrays
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[2] = "Peach"
	fruitList[3] = "Banana"
	p(fruitList[0])
	p("sier of fruitList", len(fruitList))

	var FreindList = [4]string{"Johe", "Alex", "Bob", "Michel"}
	p(FreindList[0])
	p(FreindList)
	p(len(FreindList))

	// dynamic list

	var postList = []string{"facebook", "Instagram", "twiter"}
	// [0:n-1]
	postList = append(postList[:2])

	p(postList)

	//
	highScore := make([]int, 4)
	highScore[0] = 100
	highScore[1] = 200
	highScore[2] = 300
	highScore[3] = 400
	// here append   going to act (dynamic array if size not enough)
	highScore = append(highScore, 55, 553, 434)
	sort.Ints(highScore)

	var floatList = []float64{0.23, 2.32, 0.193, 23.22}
	sort.Float64s(floatList)
	p(floatList)

}

// sort.Ints() → for int
// sort.Strings() →  for Strings
// sort.Float64s() → for floats
