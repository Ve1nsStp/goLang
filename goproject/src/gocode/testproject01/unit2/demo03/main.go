package main 
import (
	"fmt"
    "unsafe"
)
func main() {
	// 定义一个整数类型

	// var num1 int8 = -130  // 超出了int8的范围
	// var num1 int8 = 130  // 超出了int8的范围 -128——127
	var num1 int8 = 120
	fmt.Println(num1)
	
	// 定义一个 负数 用uint就会不符合类型
	// var num2 uint8 = -20  // 0——255
	var num2 uint8 = 20
	fmt.Println (num2)

	// 定义一个数字 检测其类型 不设置类型默认为int
	var num3 = -545651 
	// Printf函数的作用是：格式化的，
	fmt.Printf("num3的类型是：%T",num3)
	// 检测num占用的空间，unsafe.Sizeof
	fmt.Println()
	fmt.Println(unsafe.Sizeof(num3))
}