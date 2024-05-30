package main

import(
	// "fmt"
	// "go/ast"
	// "go/parser"
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

func lint(deref derefPointerStorage, store storage, warnStorage *warningStorage) {
	for funcName, funcDeref := range(deref) {
		
	}
}