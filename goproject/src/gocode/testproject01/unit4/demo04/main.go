package main
import "fmt"
func main(){
	// 实现一个功能，根据给出学生分数，判断学生等级
	// 给定分数
	var score int = 187
	// 判断等级
	switch score/10 {
	case 10 :
		fmt.Println("您的等级为A级")
	case 9 :
		fmt.Println("您的等级为A级")
	case 8 :
		fmt.Println("您的等级为B级")
	case 7 :
		fmt.Println("您的等级为b级")
	case 6 :
	    fmt.Println("您的等级为C级")
	case 5 :
		fmt.Println("您的等级为e级")
	case 4 :
		fmt.Println("您的等级为E级")
	case 3 :
		fmt.Println("您的等级为E级")
	case 2 :
		fmt.Println("您的等级为E级")
	case 1 :
		fmt.Println("您的等级为E级")
	case 0 :
		fmt.Println("您的等级为E级")
	default :
		fmt.Println("您录入成绩有误")
	}
}