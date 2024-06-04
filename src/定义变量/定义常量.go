package main

import "fmt"

func main() {
	const a = 100

	const b int = 100

	const (
		c = 2
		d
		e = "asd"
		f
	)

	fmt.Println(c, d, e, f)
}
