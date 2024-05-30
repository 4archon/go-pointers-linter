package main

import (
	// "fmt"
	// "go/ast"
	// "go/parser"
	"fmt"
	"go/token"
	// "os"
)

type warning struct {
	message string
	level int
	position token.Pos
	lastChangePos token.Pos
}

type warningStorage []warning

func (store *warningStorage) add(message string, level int, position token.Pos, lastChangePos token.Pos) {
	var warn warning
	
	warn.message = message
	warn.level = level
	warn.position = position
	warn.lastChangePos = lastChangePos
	
	*store = append(*store, warn)
}

func (store *warningStorage) print(fset *token.FileSet) {
	for _, i := range(*store) {
		fmt.Println(fset.Position(i.position).String())
		fmt.Println(i.level)
		fmt.Println(i.message)
		fmt.Println(fset.Position(i.lastChangePos).String())
		fmt.Println("----------")
	}
}

func lint(deref derefPointerStorage, store storage, warnStorage *warningStorage, fset *token.FileSet) {
	for funcName, funcDeref := range(deref) {
		for _, i := range(funcDeref) {
			var warn warning
			index := store.getLastFromDeref(funcName, i, fset)
			if index != -1 {
				value, pos := store.getValueAndPosVarByIndex(funcName, i.varName, index)
				switch value {
				case "nil":
					warn.level = 1
					warn.lastChangePos = pos
					warn.message = "null pointer dereference"
					warn.position = i.pos

					*warnStorage = append(*warnStorage, warn)
				
				case "func":
					warn.level = 2
					warn.lastChangePos = pos
					warn.message = "pointer might be null"
					warn.position = i.pos

					*warnStorage = append(*warnStorage, warn)
				}
			}
		}
	}
}