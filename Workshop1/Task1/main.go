package main

import (
	"fmt"
	"strconv"
)

const roundsCount int = 100

func main() {
	fmt.Print(fizzBuzz(roundsCount))
}

func fizzBuzz(roundsCount int) string {
	var result string = ""
	for i := 1; i <= roundsCount; i++ {
		if i > 1 {
			result += ", "
		}

		if i%3 == 0 {
			result += "Fizz"
		} else if i%5 == 0 {
			result += "Buzz"
		} else {
			result += strconv.Itoa(i)
		}
	}

	return result
}
