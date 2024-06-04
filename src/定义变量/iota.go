package main

import "fmt"

func main() {
	// const (
	// 	a = iota
	// 	b
	// 	c
	// 	d
	// )
	// fmt.Println(a, b, c, d)

	const (
		a = iota
		b = "hello"
		c
		d = iota
	)
	const e = iota
	fmt.Println(a, b, c, d, e)
}
