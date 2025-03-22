package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello,ninjas !")
	fmt.Println(calculatePrice(1, 2, 3))
}

func calculatePrice(x, y, z int) int {
	return x + y + z
}
