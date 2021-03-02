package main

import "fmt"

func main() {
	//TODO
	/**
	1.指针，指针的指针（c语言：二级指针）
	2.指针数组
	3.数组指针
	4.函数指针
	5.指针作为参数：引用传递，操作函数中的指针会直接改变参数的原始数据


	*/
	//	1.数组指针
	arr1 := [4]int{1, 2, 3, 4}
	var p1 *[4]int
	p1 = &arr1
	fmt.Printf("%p\n", p1)
	fmt.Println(arr1)
	fmt.Printf("%d\n", (*p1)[2]) //取数组指针的值
	fmt.Printf("%d\n", p1[1])    //取数组指针的值、简化写法
	//指针数组操作数组数据
	(*p1)[2] = 100
	fmt.Println(arr1)
	//fmt.Printf("%p\n",p1)
	//	2.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]*int{&a, &b, &c, &d}
	fmt.Println(arr2)
	fmt.Printf("%d\n", *arr2[2])
	fmt.Printf("%d\n", *arr2[2])

}
