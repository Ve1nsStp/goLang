package main
import "fmt"
func main(){
	// 与逻辑：&& 两个数值或者表达式 只要有一侧是false，结果一定是false
	// 短路与：左侧的第一个值为false 右侧的表达式或者值都不执行了
	fmt.Println(true&&true)
	fmt.Println(true&&false)
	fmt.Println(false&&true)
	fmt.Println(false&&false)
	
	// 或逻辑：|| 只要有一侧是ture 结果一定是true
	// 短路或：左侧第一个值为true 右侧的表达式或者值都不执行了
	fmt.Println(true||true)
	fmt.Println(true||false)
	fmt.Println(false||true)
	fmt.Println(false||false)
	
	// 非逻辑：! 取相反的结果
	fmt.Println(!true)
	fmt.Println(!false)
}