package main

import (
	// "errors"
	"fmt"
	"go/ast"
	// "go/parser"
	// "go/token"
	// "os"
)


type funcStorage []*ast.FuncDecl

type funcReturns map[string][]string

func (ret funcReturns) printReturns() {
	for i, j := range ret {
		fmt.Println(i)
		for _, y := range j {
			fmt.Print(y)
			fmt.Print(", ")
		}
		fmt.Println()
		fmt.Println("----------")
	}
}

func (ret *funcReturns) init() {
	*ret = make(funcReturns)
}

func (ret funcReturns) getReturns(store funcStorage) {
	for _, i := range store {
		if i.Type.Results != nil {
			var list []string
			for _, j := range i.Type.Results.List {
				switch x := j.Type.(type) {
				case *ast.StarExpr:
					list = append(list, "*" + x.X.(*ast.Ident).Name)
				case *ast.Ident:
					list = append(list, x.Name)
				}
			}
			ret[i.Name.Name] = list
		} else {
			ret[i.Name.Name] = nil
		}
	}
}

func (store *funcStorage) findFunctions(node ast.Node) bool {
	switch funcDec := node.(type) {
	case *ast.FuncDecl:
		*store = append(*store, funcDec)
	}
	return true
}