package main

import "fmt"

func foo() *int {
	return nil
}

func main() {
	var i1 int = 20
	i1 = 30
	var i6 *int
	// var i2, i4 *int = &i1, &i1
	// i2 = &i1
	i5 := &i1
	// i4 = i2
	// i3 := foo()
	// fmt.Println(i1)
	// fmt.Println(*i2)
	// fmt.Println(*i3)
	// fmt.Println(*i4)
	// fmt.Println(*i5)
	// fmt.Println(*i6)
}
