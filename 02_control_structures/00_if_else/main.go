package main

import "fmt"

func f1() int {
	x, y := 10, 5
	if x > 0 {
		return y
	}
	return 0
}

func main() {
	// simple if control
	fmt.Printf("x = %v \n", f1())
}
