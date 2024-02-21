package main
import "fmt"
func main(){
	// 实现一个功能，如果口罩的库存小于30个，提示库存不足 大于30 ，提示 库存充足
	var count int = 100
	if count <30 {
		fmt.Println("对不起，口罩数量不足")
	} else {
		fmt.Println("口罩数量充足")
	}
	// 双分支二选一，一定会选择其中一个分支
	// else 不能换到下一行
}