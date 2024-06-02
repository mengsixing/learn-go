package main

import "fmt"

var a = 1

var b int = 2

// 全局无法使用 := 定义变量
// gg:=2

func main() {
	c := false
	cc := "hello nihao"
	var dd, dd2, dd3 = 1, '2', 3
	var (
		ee = 2
		ff = false
	)
	fmt.Println("c", c, cc)
	fmt.Println(dd, dd2, dd3, ee, ff)
}
