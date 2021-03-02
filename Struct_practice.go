package main

import "fmt"

func main() {
	/**
	1.结构体是值类型

		可以通过指针取操作结构体，变成引用类型
		使用内置函数new()来创建指针（new()是go语言中专用用来声明指针变量的关键字）
		通过new创建出来的指针不是nil，而是0值
	2.声明 type 名称 struct{
			属性名字  类型
			属性名字  类型...
		}
	3.使用
	4.匿名结构体
		类似java的匿名函数
		变量名 := struct{
			member definition
	   		member definition
		}{
			name: "definition",
		   desp: "definition"

		}
	5.结构体的匿名字段
	声明 type 名称 struct{
		类型//只有类型的字段就是匿名字段，默认使用类型作为名字，匿名字段不能重复
		类型...
	}
	6.结构体嵌套
	结构体中的属性为另一个结构体，这样的结构体就是结构体嵌套
	访问嵌套结构体中的属性，变量.属性.结构体属性,例如：p1.addr.name..
	嵌套的时候依然保持值传递，当使用结构体指针作为嵌套的时候，则表现为引用类型传递
	*/
	p1 := Person{"张三", 20, "北京"} //结构体声明方式1
	fmt.Println(p1)
	fmt.Printf("%s,%d,%s\n", p1.name, p1.age, p1.addr) //结构体属性访问,p1.属性名
	p2 := Person{}                                     //结构体声明方式2
	p2.name = "sdd"                                    //结构体属性赋值
	p2.age = 23
	p2.addr = " 23"
	fmt.Println(p2)

}

type Person struct {
	name string
	age  int
	addr string
}
