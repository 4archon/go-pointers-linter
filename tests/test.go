package main

import "fmt"

func foo() *int {
	return nil
}

func foo2() (int, int) {
	return 1, 2
}

func foo3() (*int, *int) {
	return nil, nil
}

func main() {
	var i1, i2 int = foo2()
	var i3, i4 *int = foo(), foo()
	var i5, i6 *int = foo3()
	// var i1 int = 20
	// i1 = 30
	// var i6 *int
	// var i2, i4 *int = &i1, i6
	// i6 = i4
	// i2 = &i1
	// i5 := &i1
	// i4 = i2
	// i3 := foo()
	// i6 = i4
	// fmt.Println(i1)
	// fmt.Println(*i2)
	// fmt.Println(*i3)
	// fmt.Println(*i4)
	// fmt.Println(*i5)
	// fmt.Println(*i6)
}
