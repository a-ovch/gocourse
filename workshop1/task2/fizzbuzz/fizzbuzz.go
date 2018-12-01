package fizzbuzz

import (
	"bytes"
	"strconv"
)

func FizzBuzz(roundsCount int) string {
	var resBuffer bytes.Buffer
	for i := 1; i <= roundsCount; i++ {
		if i > 1 {
			resBuffer.WriteString(", ")
		}

		isFizz := i%3 == 0
		isBuzz := i%5 == 0

		if isFizz && isBuzz {
			resBuffer.WriteString("Fizz Buzz")
		} else if isFizz {
			resBuffer.WriteString("Fizz")
		} else if isBuzz {
			resBuffer.WriteString("Buzz")
		} else {
			resBuffer.WriteString(strconv.Itoa(i))
		}
	}

	return resBuffer.String()
}
