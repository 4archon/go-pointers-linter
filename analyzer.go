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
	var astPrint string;
	if len(os.Args) == 2 {
		astPrint = os.Args[1]
	}

	var scandir string = "tests/"
	fset := token.NewFileSet()


	node, err := parser.ParseDir(fset, scandir, nil, 0)
	if err != nil {
		fmt.Println("Parsing error")
		os.Exit(1)
	}

	if astPrint == "ast" {
		ast.Fprint(os.Stdout, fset, node, nil)
	} else if astPrint == "anl" {
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


	}

}
