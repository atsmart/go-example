package main

import "fmt"

/*
常规的数组声明方法

 [32]byte     // 长度为32的数组，每个元素为一个字节
 [2*N] struct { x, y int32 } // 复杂类型数组
 [1000]*float64   // 指针数组
 [3][5]int    // 二维数组
 [2][2][2]float64   // 等同于[2]([2]([2]float64))

 数组长度在定义后就不可更改，在声明时长度可以为一个常量或者一个常量 表达式（常量表达式是指在编译期即可计算结果的表达式）。
 数组的长度是该数组类型的一个内 置常量，可以用Go语言的内置函数len()来获取。下面是一个获取数组arr元素个数的写法：
 arrLength := len(arr)

*/

var (
	x, y = 1, 2
)

/*
 **IMPORT**
 在Go语言中数组是一个值类型（value type）。所有的值类型变量在赋值和作为参数传递时都将产生一次复制动作。
 如果将数组作为函数的参数类型，则在函数调用时该 参数将发生数据复制。因此，在函数体中无法修改传入的数组的内容，
 因为函数内操作的只是所 传入数组的一个副本。
*/
func Modify(array [5]int) {
	array[0] = 10
	fmt.Printf("In Modify(), array values: %v \n", array)
}

func main() {
	arr := [3]int{1, 2, 3}
	for i, v := range arr {
		fmt.Printf("index = %v, value = %v \n", i, v)
	}

	//索引器的访问数组方式
	fmt.Printf("arr[0] = %v\n", arr[0])

	// 由程序自动推算出数据的长度，可以使用...来代表数据的长度
	// 第9个元素赋值为1,其他元素为默认值0
	arr2 := [...]int{9: 10}
	fmt.Printf("arr2 length = %v\n", len(arr2))
	sum := 0
	for _, v := range arr2 {
		fmt.Printf(", %v", v)
		sum += v
	}
	fmt.Printf("\nsum = %v \n", sum)

	/*
		在go语言中，数组长度也是做为结构的一部分，[1]int 和 [2]int 是两种不同的数据类型
	*/
	//	数据长度len 等于10， 容量cap 等于20
	arr3 := make([]int, 10, 20)
	index := 0
	for i, _ := range arr3 {
		arr3[i] = index
		index++
	}
	sum = 0
	for _, v := range arr3 {
		sum += v
	}
	fmt.Printf("arr3 sum = %v\n", sum)

	//指针数组，里面存放的数据是指针
	arr4 := [...]*int{&x, &y}
	for _, v := range arr4 {
		fmt.Printf(" ,%v", v)
	}
	fmt.Println()
	//数组指针
	var parr *[len(arr)]int = &arr
	fmt.Printf("point parr = %v, type = %T \n", parr, parr)

	for _, v := range *parr {
		fmt.Printf(" ,%v", v)
	}
	fmt.Println()

	parr2 := new([len(arr)]int)
	parr2 = &arr
	fmt.Printf("point parr2 = %v ,type = %T \n", parr2, parr2)

	//二维数组初始化的方式
	arr5 := [...][2]int{
		{1, 2},
		{3, 4}}
	for _, v := range arr5 {
		for _, v1 := range v {
			fmt.Printf("%v, ", v1)
		}
	}
	fmt.Println()
	arr6 := [2][3]int{
		{1, 1, 1},
		{2, 2, 2}}
	fmt.Printf("arr6 = %v \n", arr6)

	array := [...]int{1, 2, 3, 4, 5}
	Modify(array)
	fmt.Printf("In main(), array values: %v \n", array)

}
