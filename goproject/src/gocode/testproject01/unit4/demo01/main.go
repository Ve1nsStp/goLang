package main

import "fmt"

func main(){
		// 实现一个功能，如果口罩的库存小于30个，提示库存不足
		var count int = 100
		if count <30 {
			fmt.Println("对不起，口罩数量不足")
		}
		// if 后面的表达式，返回结果一定是true 或者 false
		// 返回为true｛｝执行，返回为false｛｝不执行
		// if 后面一定要有空格，和条件表达式分隔开来
		// ｛｝一定不能省略
		// 表达式左右的（）是建议省略的
		// 在golang中，if后面的可以并列的加入变量的定义 用;隔开
		if num:=20;num <30 {
			fmt.Println("今天天气真不错啊")
		}
}