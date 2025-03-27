package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	fmt.Println("Good morning " + n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v \n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("mario")
	sayBye("luigi")
	cycleNames([]string{"luigi", "mario", "crystal"}, sayGreeting)
	a1 := circleArea(10)
	fmt.Println("circle is ", a1)
}
