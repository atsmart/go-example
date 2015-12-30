package main

import (
	"fmt"
	//	"strconv"
)

func main() {
	//	字符串的内容可以用类似于数组下标的方式获取，但与数组不同，字符串的内容不能在初始 化后被修改
	str := "hello golang"
	size := len(str)
	var ch byte
	for i := 0; i < size; i++ {
		ch = str[i]
		fmt.Printf("_%v", ch)
	}
	fmt.Println("")

	for i, v := range str {
		fmt.Printf("index = %v, value = %v \n", i, v)
	}
}
