package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name:", reader)

	b := newBill(name)
	fmt.Println("created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("choose option ( a - add item, s - save bill, t - add tip):", reader)
	fmt.Println(opt)
}

func main() {
	mybill := createBill()
	promptOptions(mybill)
}
