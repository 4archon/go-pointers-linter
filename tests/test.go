package main

import "fmt"

func foo() int {
	var i2 int = 100
	return i2
}

func main() {
	var i1 int = 20
	i1 = 30
	i1++
	i1--
	i1 = i1 + 100
	i1 += 22
	i1 = foo()
	var i2 *int
	i2 = &i1
	fmt.Println(i1)
	fmt.Println(*i2)
}