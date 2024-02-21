package main
import "fmt"
var (
	n8 = 100
	n9 = "hello"
)
func main(){
	// 1. 第一种：指定变量类型，并且赋值
	var num int = 10
	fmt.Println(num)
	// 2. 第二种：指定变量类型，声明后若不赋值，使用默认值
	var num2 int 
	fmt.Println(num2) // int类型的默认值为0
	// 3. 第三种：根据值自行判定变量类型（类型推导）
	var num3 = 10.11
	fmt.Println(num3)
	// 4. 第四种：省略var，注意 := 左侧的变量不应该是已经声明过的，否则会导致编译错误
	// name := "赵云"
	fmt.Println("____________________分割线_______________________")
	// 一次性声明多个变量
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)
	// 声明多个变量并赋值
	var n4, name1, n5 = 100, "张飞", 888
	fmt.Println(n4, name1, n5)
	// 使用海豹运算符声明多个变量:=
	n6, name2, n7 := 100, "关羽", 888 
	fmt.Println(n6, name2, n7)
	
	fmt.Println(n8, n9)
	}	