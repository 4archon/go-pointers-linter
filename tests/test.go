package main

import "fmt"

func foo() *int {
	return nil
}

func main() {
	var i1 int = 20
	i1 = 30
	var i2 *int
	i2 = &i1
	i3 := foo()
	fmt.Println(i1)
	fmt.Println(*i2)
	fmt.Println(*i3)
}
