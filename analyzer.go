package main

import (
	// "errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)




func main() {
	var command string
	var dir string
	if len(os.Args) == 3 {
		dir = os.Args[1]
		command = os.Args[2]
	} else {
		fmt.Println("run: analyzer [scandir] command[ast|run]")
		os.Exit(2)
	}

	var scandir string = dir
	fset := token.NewFileSet()


	node, err := parser.ParseDir(fset, scandir, nil, 0)
	if err != nil {
		fmt.Println("Parsing error")
		os.Exit(1)
	}

	if command == "ast" {
		ast.Fprint(os.Stdout, fset, node, nil)
	} else if command == "anl" {
		var funcStore funcStorage

		var s storage
		s.init()

		ast.Inspect(node["main"], funcStore.findFunctions)

		var ret funcReturns
		ret.init()

		ret.getReturns(funcStore)
		// ret.printReturns()

		analyzeFunc(funcStore, s, ret)
		// s.printStore(fset)

		var deref derefPointerStorage
		deref.init()

		getFuncStarIden(funcStore, deref, s)
		// deref.print(fset)

		var warn warningStorage

		lint(deref, s, &warn, fset)

		// warn.print(fset)

		printer(warn, fset)
	} else if command == "run" {
		var funcStore funcStorage

		var s storage
		s.init()

		ast.Inspect(node["main"], funcStore.findFunctions)

		var ret funcReturns
		ret.init()

		ret.getReturns(funcStore)

		analyzeFunc(funcStore, s, ret)

		var deref derefPointerStorage
		deref.init()

		getFuncStarIden(funcStore, deref, s)

		var warn warningStorage

		lint(deref, s, &warn, fset)

		printer(warn, fset)
	}

}
