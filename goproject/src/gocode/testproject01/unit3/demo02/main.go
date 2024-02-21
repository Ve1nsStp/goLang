package main

import "fmt"

func main(){
	var num1 int = 10
	fmt.Println(num1)
	const num2 int = (10+30)%3 +3 -7 // 等号右侧运算结束后赋给左侧
	fmt.Println(num2)

	var num3 int = 10 
	num3 += 20 // num3 = num3 +20
	fmt.Println(num3)
	// 练习：交换两个数的值并输出结果
	var a int = 8 
	var b int = 4
	fmt.Println("a=%v,b=%v",a,b)
	// 引入第三变量
	// var temp int 
	// temp = a 
	// a = b
	// b=temp
	// fmt.Println("a=%v,b=%v",a,b)
	// 直接交换
	a,b = b,a 
	fmt.Println("a=%v,b=%v",a,b)
}