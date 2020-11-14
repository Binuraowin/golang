package main

import "fmt"

var i float32 = 42
var (
	actorName string = "binura"
	drama     string = "bites"
	dramaNum  int    = 3
	season    int    = 11
)

func main() {
	x := 5
	y := 6
	var sum = x + y
	fmt.Println(sum)
	if sum > 15 {
		fmt.Println("ok")
	} else if sum < 10 {
		fmt.Println("try 1")
	} else {
		fmt.Println("try 2")
		fmt.Printf("%v, %T", x, y)
	}

	var a [5]int

	fmt.Println(a)
	fmt.Printf("%v, %T", i, i) //this works only in priintf
}
