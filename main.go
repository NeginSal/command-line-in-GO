package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

	mybill.addItem("onion soup", 4.50)
	mybill.addItem("veg pie", 5.35)
	mybill.addItem("pudding", 3.25)
	mybill.addItem("tea", 7.00)
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
