package main

import "fmt"

// 常量定义
const PI float64 = 3.14159265358979323846
const ZERO = 0.0 // 无类型浮点常量

const (
	SIZE int64 = 1024
	EOF        = -1 // 无类型整型常量
)

const u, v float32 = 0, 3 // u = 0.0, v = 3.0，常量的多重赋值
const MASK = 1 << 3

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays // 这个常量没有导出
)

// 变量定义
var v1 int
var v2 string
var v3 [10]int // 数组
var v4 []int   // 数组切片
var v5 struct{ f int }
var v6 *int           // 指针
var v7 map[string]int // map，key为string类型，value为int类型
var v8 func(a int) int

//全局变量组
var (
	v11 int
	v21 string
)

func Name() (firstName, lastName, nickName string) {
	return "Eric", "Shang", "Eric.Shag"
}

func main() {
	v1, v11 = 5, 10
	//	多重赋值,在不支持多重赋值的语言中，交互两个变量的内容需要引入一个中间变量：
	// t = i; i = j; j = t;
	v1, v11 = v11, v1
	fmt.Printf("v1 = %v, v11=%v \n", v1, v11)

	//匿名变量
	_, _, nickName := Name()
	fmt.Printf("nickName = %v\n", nickName)

	fmt.Printf("Const Mask = %v \n", MASK)
}
