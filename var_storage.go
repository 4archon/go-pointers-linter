package main

import (
	// "errors"
	"fmt"
	"go/ast"
	// "go/parser"
	"go/token"
	// "os"
)

type varValue struct {
	value string
	pos token.Pos
}

type varStore struct {
	name string
	typeVar string
	last int
	listVar []varValue
}

type storage map[string]map[string]varStore

func (store *storage) init() {
	*store = make(storage)
}

func (store storage) init2lvl(str string) {
	store[str] = make(map[string]varStore)
}

func getType(node *ast.StarExpr) (string) {
	switch x := node.X.(type) {
	case *ast.Ident:
		return "*" + x.Name
	}
	return ""
}

func (store storage) addStarExpr (node *ast.ValueSpec, expr *ast.StarExpr, funcName string) {
	var s varStore
	s.name = node.Names[0].Name
	s.typeVar = getType(expr)
	s.last = 0

	var v varValue
	v.pos = node.Names[0].NamePos
	v.value = "nil"

	s.listVar = append(store[funcName][s.name].listVar, v)

	store[funcName][s.name] = s
}

// func (store storage) addExist ()

func (store storage) printStore(fset *token.FileSet) {
	for _, i := range store {
		for _, j := range i {
			for _, k := range j.listVar {
				fmt.Println(j.name)
				fmt.Println(getLine(fset, k.pos))
				fmt.Println(j.typeVar)
				fmt.Println(k.value)
				fmt.Println("--------")
			}
		}
	}
}