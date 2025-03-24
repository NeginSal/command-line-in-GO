package main

import (
	"fmt"
)

func main() {
	// Arrays
	var ages [3]int = [3]int{20, 30, 40}
	var agesTwo [2]int = [2]int{50, 60}

	fmt.Println("ages=", ages, "length=", len(ages))
	fmt.Println("agesTwo=", agesTwo, "length=", len(agesTwo))

	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[1] = "luigi"
	fmt.Println("names", names, len(names))

	// Slices
	var scores = []int{100, 50, 60}
	scores[2] = 25
	fmt.Println("scores", scores)

	scores = append(scores, 85)
	fmt.Println("scores", scores)

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
