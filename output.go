package main

import (
	"fmt"
	"go/token"
)

func printer (warn warningStorage, fset *token.FileSet) {
	l1 := "Error"
	l2 := "Warning"
	for _, i := range(warn) {
		warnPos := fset.Position(i.position).String()
		changePos := fset.Position(i.lastChangePos).String()
		if i.level == 1 {
			fmt.Printf("%s: ", l1)
		} else if i.level == 2 {
			fmt.Printf("%s: ", l2)
		}
		fmt.Printf("%s ", i.message)
		fmt.Printf("at %s\n", warnPos)
		fmt.Printf("Last change at %s\n", changePos)
		fmt.Println("------------")
	}
}