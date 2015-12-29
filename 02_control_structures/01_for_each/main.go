package main

import (
	"fmt"
)

/*

// Like a C for
for init; condition; post { }

// Like a C while <=> while() {}
for condition { }

// Like a C for(;;) <=> do{}while()
for { }

*/

//for init; condition; post { }
func f1() int {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

// Like a C while
func f2() int {
	m := map[string]int{"1": 1, "2": 2, "3": 3}
	sum := 0
	for k, v := range m {
		fmt.Printf("f2 key = %v, value = %v \n", k, v)
		sum += v
	}

	array := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range array {
		sum += v
	}
	return sum
}

func f3() string {
	i, sum := 0, 0
	for i < 10 {
		sum += i
		i++
	}
	return fmt.Sprintf("f3 sum = %v", sum)
}

func main() {
	fmt.Printf("f1 sum = %v \n", f1())
	fmt.Printf("f2 sum = %v \n", f2())
	fmt.Println(f3())
}
