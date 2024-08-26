package main

import "fmt"

func main() { /*{不能单独一行*/
	var a int = 1
	fmt.Println(a)
	var s string = fmt.Sprintf("%dHello", a)
	fmt.Println(s + ", World!")
	var b = "zkh"
	fmt.Println(b)
	c := 2
	fmt.Println(c)
	d := "mfy"
	fmt.Println(d)
	var e float32
	e = 3.14
	fmt.Println(e)
	b, d = d, b
	fmt.Println(b)
	/*函数题内部才可以:=初始化声明，声明或者声明完赋值的变量必须使用*/
	/*字符串拼接数字使用方式类似c中print输出函数
	printf是直接输出，sprintf是先返回一个字符串*/
	// var indetifier []int
	// indetifier = make([]int, 5)

}
