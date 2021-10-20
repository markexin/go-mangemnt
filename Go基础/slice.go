package main

import "fmt"

func main() {

	// 定义数组
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	// 定义切片slice
	sa := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sa)

	// 使用make定义切片
	sa2 := a[:]
	fmt.Println(sa2)

	//指针取值
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	// 指针小练习
	// 程序定义一个int变量num的地址并打印
	// 将num的地址赋给指针ptr，并通过ptr去修改num的值

	num := 10
	pNum := &num
	// fmt.Print(pNum)
	ptr := pNum
	// fmt.Print(*ptr)
	*ptr = 20
	fmt.Print(num)
}
