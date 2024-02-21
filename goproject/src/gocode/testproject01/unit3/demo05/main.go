package main
import "fmt"
func main(){
	var age = 18
	fmt.Println("age对应的存储空间地址为",&age) //age对应的存储空间地址
	var ptr *int = &age
	fmt.Println(ptr)
	fmt.Println("ptr这个指针指向的具体数值为",*ptr)
}