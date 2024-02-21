package main 
import "fmt"
func main (){
	// 定义浮点类型的数据
	var num1 float32 = 3.14
	fmt.Println(num1)
	// 可以表示正浮点数，也可以表示负的浮点数
	var num2 float32 = -3.14
	fmt.Println(num2)
	// 可以用十进制来表示，也可以用科学计数法来表示 E/e 大小写都可以
	// 科学计数法
	var num3 float32 = 314E-2 // 314乘以10的-2次方
	fmt.Println(num3)
	var num4 float32 = 314E+2 // 314乘以10的+2次方
	fmt.Println(num4)
	var num5 float32 = 314e+2 // 314乘以10的+2次方
	fmt.Println(num5)
	var num6 float64 = 314E+2 // 314乘以10的+2次方
	fmt.Println(num6)
	// 浮点数可能会有精度的丢失，所以通常情况下，建议使用float64
	var num7 float32 = 256.000000916
	fmt.Println(num7)
	var num8 float64 = 256.000000916
	fmt.Println(num8)
	// 在golang中，默认的浮点类型是float64
	var num9 = 3.156
	fmt.Printf("num9的类型是%T",num9)
}