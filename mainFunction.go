package main

import (
	"fmt"
)

// creates Greetings function
func Greetings(n string) {
	fmt.Printf("Hey there good morning %v \n", n)
}

// creates Goodbye function
func Goodbye(n string) {
	fmt.Printf("That's all folks: %v  (^^)/ \n", n)
}

func main() {
	Greetings("Cloud")
	Greetings("Tifa")
	Greetings("Barret")
	Goodbye("Aerith")
	Goodbye("vincent")
}
