package main

import (
	// "errors"
	// "fmt"
	"go/ast"
	// "go/parser"
	// "go/token"
	// "os"
)


type funcStorage []*ast.FuncDecl

func (store *funcStorage)findFunctions(node ast.Node) bool {
	switch funcDec := node.(type) {
	case *ast.FuncDecl:
		*store = append(*store, funcDec)
	}
	return true
}

func analizeFunc(funcStore funcStorage, store storage) {
	for _, i := range(funcStore) {
		store.analizeFuncBody(i.Body, i.Name.Name)
	}
}