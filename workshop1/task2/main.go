package main

import (
	"fmt"
	"github.com/a-ovch/gocourse/workshop1/task2/fizzbuzz"
)

const roundsCount int = 100

func main() {
	fmt.Printf(fizzbuzz.FizzBuzz(roundsCount))
}
