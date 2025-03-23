package main

import (
	"fmt"
)

func main() {
	age := 31
	name := "negin"

	// Print
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("newline \n")

	//Println
	fmt.Println("hello ninjas!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my name is", name, "and my age is ", age)

	//Printf(formatted strings)
	fmt.Printf("my name is %v and my age is %v \n", name, age)
	fmt.Printf("my name is %q and my age is %q \n", name, age)
	fmt.Printf("age is type of %T", age)
}
