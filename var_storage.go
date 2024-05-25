package main

import (
	// "errors"
	"fmt"
	// "go/ast"
	// "go/parser"
	"go/token"
	// "os"
)

type varProg struct {
	name string
	typeVar string
	value *int
	pos token.Pos
}

type storage map[string]map[string][]varProg

func (store storage) printStore(fset *token.FileSet) {
	for _, i := range store {
		for _, j := range i {
			for _, k := range j {
				fmt.Println(k.name)
				fmt.Println(getLine(fset, k.pos))
				fmt.Println(k.typeVar)
				fmt.Println(k.value)
				fmt.Println("--------")
			}
		}
	}
}