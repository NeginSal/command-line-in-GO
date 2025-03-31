package main

import "fmt"

func updateName(x *string) {
	*x = "wedgge"
}

func main() {
	name := "tifa"

	// updateName(name)
	// fmt.Println("memory address of name is:", &name)

	m := &name
	// fmt.Println("memory address:", m)
	// fmt.Println("value at memory address:", *m)

	fmt.Println(name)
	updateName(m) 
	fmt.Println(m)
	fmt.Println(name)
}
