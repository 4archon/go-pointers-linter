package main

import (
	// "errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func getType(node *ast.StarExpr) (string) {
	switch x := node.X.(type) {
	case *ast.Ident:
		return "*" + x.Name
	}
	return ""
}

func (store storage) analizeDeclStmt(node *ast.GenDecl, funcName string) {
	for _, i := range node.Specs {
		switch x := i.(type) {
		case *ast.ValueSpec:
			switch y := x.Type.(type) {
			case *ast.StarExpr:
				var s varProg
				s.name = x.Names[0].Name
				s.typeVar = getType(y)
				s.pos = x.Names[0].NamePos
				s.value = nil
				store[funcName][x.Names[0].Name] = append(store[funcName][x.Names[0].Name], s)
			}
		}
	}
}

func (store storage) analizeFuncBody(node *ast.BlockStmt, funcName string) {
	store[funcName] = make(map[string][]varProg)
	for _, i := range node.List {
		switch x := i.(type) {
		case *ast.DeclStmt:
			switch y := x.Decl.(type) {
			case *ast.GenDecl:
				store.analizeDeclStmt(y, funcName)
			}
		}
	}
}

func getLine(fset *token.FileSet, pos token.Pos) int {
	return fset.Position(pos).Line
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
		s := make(storage)
		ast.Inspect(node["main"], funcStore.findFunctions)
		analizeFunc(funcStore, s)
		s.printStore(fset)
	}

}
