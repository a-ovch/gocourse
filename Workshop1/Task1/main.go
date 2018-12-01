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

		isFizz := i%3 == 0
		isBuzz := i%5 == 0

		if !isFizz && !isBuzz {
			result += strconv.Itoa(i)
			continue
		}

		if isFizz {
			result += "Fizz"

			if isBuzz {
				result += " "
			}
		}

		if isBuzz {
			result += "Buzz"
		}
	}

	return result
}
