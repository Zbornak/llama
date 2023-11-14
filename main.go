/*
CS50
Lab 1: Population Growth
Author: zbornak
Date: 14/11/2023
*/

package main

import (
	"fmt"
)

func main() {
	var initialCount int
	var endingCount int

	fmt.Println(" '")
	fmt.Println("~)")
	fmt.Println(" (_---;)")
	fmt.Println("  /|~|\\")
	fmt.Println(" / / / | *** LLAMA GROWTH CALCULATOR ***")

	// user input
	for initialCount < 9 {
		fmt.Printf("Enter starting population size (minimum 9): ")
		fmt.Scan(&initialCount)
	}

	for endingCount < initialCount {
		fmt.Printf("Enter ending population size: ")
		fmt.Scan(&endingCount)
	}
	// end user input

	result, singplu := calculateTime(initialCount, endingCount)

	fmt.Printf("It will take %d %s to grow your LLama population to %d.\n", result, singplu, endingCount)

	fmt.Println("*** THANK-YOU FOR USING LGC ***")
}

// calculate number of years required and also return 'year/years' depending on result
func calculateTime(initial, ending int) (int, string) {
	var yearsPassed int

	for ending > initial {
		// number of births minus number of deaths
		initial += (initial / 3) - (initial / 4)
		yearsPassed++
	}

	if yearsPassed == 1 {
		return yearsPassed, "year"
	} else {
		return yearsPassed, "years"
	}
}
