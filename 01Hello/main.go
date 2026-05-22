package main

import "fmt"

// Global variable
var jwtToken = 300000

// Constant
const LoginToken string = "secret"

// Check user login status
func checkLoginStatus(isLoggedIn bool) {
	if isLoggedIn {
		fmt.Println("Logged in")
		fmt.Println("Redirecting to login page")
	} else {
		fmt.Println("Not logged in")
	}
}

// Print uint8 value
func printSmallVariable(number uint8) {
	fmt.Println("This is a small variable:", number)
}

// Print float value
func printFloatValue(val float32) {
	fmt.Println(val)
}

func main() {

	var username string = "Ghaith"
	var smallVariable uint8 = 255
	var isLoggedIn bool = true
	var floatNumber float32 = 255.43434

	printFloatValue(floatNumber)
	checkLoginStatus(isLoggedIn)
	printSmallVariable(smallVariable)

	fmt.Println(username)
	fmt.Println(isLoggedIn)

	var website = "freeCodingCamp.in"
	fmt.Println(website)

	numberOfUsers := 30000.0
	fmt.Println(numberOfUsers)
	fmt.Println(LoginToken)
}
