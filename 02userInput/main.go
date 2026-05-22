package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()

	date := fmt.Sprintf("%d-%02d-%02d-%02d-%02d",
		now.Year(),
		now.Month(),
		now.Day(),
		now.Minute(),
		now.Second(),
	)
	time.Sleep(10000 * time.Millisecond)
	p(date)

}
