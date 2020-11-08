package main

import "fmt"

func main() {
	x := 5
	y := 6
	var sum = x + y
	fmt.Println(sum)
	if sum > 10 {
		fmt.Println("ok")
	} else if sum < 12 {
		fmt.Println("try 1")
	} else {
		fmt.Println("try 2")
	}

	var a [5]int
	fmt.Println(a)
}
