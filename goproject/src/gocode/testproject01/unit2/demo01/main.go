package main 
import "fmt"
func main(){
	// 1. 声明变量
	var name string = "赵云"
	var age int
	// 2. 赋值变量
	age = 16
	
	// 3. 使用变量
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
// 	var age int = 18    重复声明会报错
// 	fmt.Println("age = ", age)
// 	# command-line-arguments
// .\main.go:13:6: age redeclared in this block
//         .\main.go:6:6: other declaration of age
     var num int = 165.566 // 声明变量的类型和赋值的类型不一致会报错
	 fmt.Println("num = ", num)
	//  .\main.go:18:20: cannot use 165.566 (untyped float constant) as int value in variable declaration (truncated)
}