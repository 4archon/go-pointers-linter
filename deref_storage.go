package main

import (
	// "errors"
	"fmt"
	"go/ast"
	// "go/parser"
	"go/token"
	// "os"
)

type AstStarExprIden []*ast.Ident

func (store *AstStarExprIden) getAstStarIden(node ast.Node) bool {
	switch x := node.(type) {
	case *ast.StarExpr:
		*store = append(*store, x.X.(*ast.Ident))
	}
	return true
}

func (store AstStarExprIden) print(fset *token.FileSet) {
	for _, i := range store {
		fmt.Println(i.Name)
		fmt.Println(getLine(fset, i.NamePos))
		fmt.Println("--------")
	}
}

func getFuncStarIden(funcStore funcStorage, storePointers derefPointerStorage, store storage) {
	for _, i := range(funcStore) {
		var starIden AstStarExprIden
		funcName := i.Name.Name
		ast.Inspect(i.Body, starIden.getAstStarIden)
		storePointers.getDeref(funcName, starIden, store)
	}
}

type derefPointer struct {
	varName string
	pos token.Pos
}

type derefPointerStorage map[string][]derefPointer

func (store *derefPointerStorage) init() {
	*store = make(derefPointerStorage)
}

func (store derefPointerStorage) print(fset *token.FileSet) {
	for k, i := range(store) {
		fmt.Println(k)
		fmt.Println("-------------")
		for _, j := range(i) {
			fmt.Println(j.varName)
			fmt.Println(getLine(fset, j.pos))
			fmt.Println("------")
		}
	}
}

func (storePointers derefPointerStorage) getDeref(funcName string, idents AstStarExprIden, store storage) {
	var list []derefPointer
	for _, i := range(idents) {
		_, isValid := store[funcName][i.Name]
		if isValid {
			var deref derefPointer
			deref.varName = i.Name
			deref.pos = i.NamePos
			list = append(list, deref)
		}
	}
	storePointers[funcName] = list
}