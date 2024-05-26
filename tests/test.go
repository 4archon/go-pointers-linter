package main

import "fmt"

func foo() *int {
	return nil
}

func main() {
	var i1 int = 20
	i1 = 30
	var i2 *int
	var i4 *int
	i2 = &i1
	i4 = i2
	i3 := foo()
	fmt.Println(i1)
	fmt.Println(*i2)
	fmt.Println(*i3)
	fmt.Println(*i4)
}
