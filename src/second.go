package main

import "fmt"

func main() {
	// var n bool = false
	// fmt.Printf("%v, %T", n, n)

	a := 10
	b := 4
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b) // oppesite of or

	fmt.Println(a << 3)
	fmt.Println(a >> 5)
}
