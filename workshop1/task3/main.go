package main

import (
	"bytes"
	"fmt"
	"github.com/a-ovch/gocourse/workshop1/task3/duck"
)

func main() {
	fmt.Println(playWithDuck(duck.NewDeafMuteDuck()))
	fmt.Println(playWithDuck(duck.NewReactiveLoudDuck()))
	fmt.Println(playWithDuck(duck.NewBackwardsQuietDuck()))
}

func playWithDuck(d duck.Duck) string {
	var buffer bytes.Buffer

	for i := 0; i < 3; i++ {
		buffer.WriteString(d.Fly())
		buffer.WriteString("||")
		buffer.WriteString(d.Quack())
		buffer.WriteString("\n")
	}

	return buffer.String()
}
