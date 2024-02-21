package main

import "fmt"

func main(){
	// + 加号
	// 1、正数操作 2、相加操作 3、字符串拼接
	var n1 int = +10
	fmt.Println(n1)
	var n2 int = 5 + 6
	fmt.Println(n2)
	var s1  = "Abd"+ "sssde"
	fmt.Println(s1)

	// / 除号
	fmt.Println(10/3) // 两个int参与运算，结果一定为整数类型
	fmt.Println(10.0/3) // 浮点类型参与运算，结果为浮点类型
	fmt.Println(10/-3) // -3
	// % 取模  等价公式 a%b=a-(a/b:此处取出整数)*b
	fmt.Println(10%3) // 1
	fmt.Println(-10%3) // -10 - (-3*3) = -1
	fmt.Println(10%-3) // 10 - (10/-3)*3 = 1
	fmt.Println(-10%-3) // -10 -(3*-3) = -1
	// ++ 自增符号
	var a int = 10 
	a++
	fmt.Println(a)
	// -- 自减符号
	a--
	fmt.Println(a)
	// golang中，++，--，只能单独使用，不能参与到运算中
	// golang中，++，--，只能跟在变量的后面，不能像js那样写在变量前面
}