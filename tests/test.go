package main

import "fmt"

// func foo() *int {
// 	var j1 string = "ewqewq"
// 	var j2 *string = nil
// 	j2 = &j1
// 	j2 = nil
// 	fmt.Println(*j2)
// 	return nil
// }

// func foo2() (int, int) {
// 	return 1, 2
// }

// func foo3() (*int, *int) {
// 	return nil, nil
// }

func main() {
	var i1 int = 30
	var p1 *int = &i1
	*p1 = 20
	p1 = &i1
	fmt.Println(*p1)

	// var j1 string = "ewqewq"
	// var j2 *string = &j1
	// j3 := j2
	// var j4 *string
	// var i1 int = 20
	// i1 = 30
	// var i6 *int
	// var i2, i4 *int = &i1, i6
	// i6 = i4
	// i2 = &i1
	// i5 := &i1
	// i4 = i2
	// i3 := foo()
	// i10 := i3
	// i6 = i4
	// fmt.Println(i1)
	// fmt.Println(*i2)
	// fmt.Println(*i3)
	// fmt.Println(*i4)
	// fmt.Println(*i5)
	// fmt.Println(*i6)
	// fmt.Println(*i10)
	// fmt.Println(j1)
	// fmt.Println(*j2)
	// fmt.Println(*j3)
	// fmt.Println(*j4)

	// var i1, i2 int = foo2()
	// var i3, i4 *int = foo(), foo()
	// var i5, i6 *int = foo3()
	// i5, i6 = foo3()
	// i7, i8 := foo3()
	// i9 := foo()
	// i7 = foo()
}
