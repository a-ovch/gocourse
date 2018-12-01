package fizzbuzz

import "strconv"

func FizzBuzz(roundsCount int) string {
	var result string = ""
	for i := 1; i <= roundsCount; i++ {
		if i > 1 {
			result += ", "
		}

		isFizz := i%3 == 0
		isBuzz := i%5 == 0

		if isFizz && isBuzz {
			result += "Fizz Buzz"
		} else if isFizz {
			result += "Fizz"
		} else if isBuzz {
			result += "Buzz"
		} else {
			result += strconv.Itoa(i)
		}
	}

	return result
}
