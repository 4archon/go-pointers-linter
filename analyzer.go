package main

import (
	// "errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func analizeDeclStmt(node *ast.GenDecl, funcName string, store storage) {
	for _, i := range node.Specs {
		switch x := i.(type) {
		case *ast.ValueSpec:
			switch y := x.Type.(type) {
			case *ast.StarExpr:
				store.addStarExpr(x, y, funcName)
			}
		}
	}
}

func analizeAssignStmt(node *ast.AssignStmt, funcName string, store storage) {
	for _, i := range node.Rhs {
		switch x := i.(type) {
		case *ast.UnaryExpr:
			if x.Op.String() == "&" {
				
			}
		case *ast.Ident:

		}
	}
}

func analizeFuncBody(node *ast.BlockStmt, funcName string, store storage) {
	store.init2lvl(funcName)
	for _, i := range node.List {
		switch x := i.(type) {
		case *ast.DeclStmt:
			switch y := x.Decl.(type) {
			case *ast.GenDecl:
				analizeDeclStmt(y, funcName, store)
			}
		case *ast.AssignStmt:

		}
	}
}

func analizeFunc(funcStore funcStorage, store storage) {
	for _, i := range(funcStore) {
		analizeFuncBody(i.Body, i.Name.Name, store)
	}
}


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
		analizeFunc(funcStore, s)
		s.printStore(fset)
	}

}
