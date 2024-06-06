package main
import "fmt"
func main(){
    // 实现一个功能，求和
//     var i1 int = 1
//     var i2 int = 2
//     var i3 int = 3
//     var i4 int = 4
//     var i5 int = 5
//        var i int = 1
    // 定义一个变量接收这个和
    var sum int = 0
    for i := 1 ; i < 6 ; i++{
    sum += i
    }
    // 输出结果
    fmt.Println(sum)
}