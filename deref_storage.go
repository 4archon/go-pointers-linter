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

func getFuncStarIden(funcStore funcStorage) {
	var starIden AstStarExprIden
	for _, i := range(funcStore) {
		funcName := i.Name.Name
		ast.Inspect(i.Body, starIden.getAstStarIden)

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

func (store derefPointerStorage) getDeref(funcName string, idents AstStarExprIden) {
	
}