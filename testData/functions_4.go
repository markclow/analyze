package main

import "fmt"

// Update marsYear so that it takes earthYears
// As a parameter
func computeMarsYears(earthYears int) int {
	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func another() int {
	return computeMarsYears(25)
}

func main() {
	myAge := 25

	// Call `marsYear` with `myAge`
	myMartianAge := computeMarsYears(computeMarsYears(myAge))
	fmt.Println(myMartianAge)
}
