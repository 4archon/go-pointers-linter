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

func analizeValueSpec(node *ast.Ident, expr *ast.StarExpr, funcName string, store storage, val string) {
	name := node.Name
	typeVar := getType(expr)
	value := val
	pos := node.NamePos

	store.addNewVar(funcName, name, typeVar, value, pos)
}

func analizeSpecValues(node *ast.Ident, expr ast.Expr, typeExpr *ast.StarExpr, funcName string, store storage) {
	var value string
	switch x := expr.(type) {
	case *ast.UnaryExpr:
		if x.Op.String() == "&" {
			value = "valid"
		}
	case *ast.Ident:
		value = store.getLastValue(funcName, x.Name)
	}
	analizeValueSpec(node, typeExpr, funcName, store, value)
}

func analizeDeclStmt(node *ast.GenDecl, funcName string, store storage) {
	for _, i := range node.Specs {
		switch x := i.(type) {
		case *ast.ValueSpec:
			if x.Values == nil {
				switch y := x.Type.(type) {
				case *ast.StarExpr:
					for _, j := range x.Names {
						analizeValueSpec(j, y, funcName, store, "nil")
					}
				}
			} else {
				//add assign stmt
				switch y := x.Type.(type) {
				case *ast.StarExpr:
					for z, j := range x.Names {
						analizeSpecValues(j, x.Values[z], y, funcName, store)
					}
				}
			}
		}
	}
}

func getIdentName(node ast.Expr) string {
	var res string
	switch x := node.(type) {
	case *ast.Ident:
		res = x.Name
	}
	return res
}

func getIdentPos(node ast.Expr) token.Pos {
	var res token.Pos
	switch x := node.(type) {
	case *ast.Ident:
		res = x.NamePos
	}
	return res
}

func analizeAssignStmt(node *ast.AssignStmt, funcName string, store storage) {
	for j, i := range node.Rhs {
		lVarPos := getIdentPos(node.Lhs[j])
		lVarName := getIdentName(node.Lhs[j])
		rVarName := getIdentName(i)
		switch x := i.(type) {
		case *ast.UnaryExpr:
			if node.Tok.String() == "=" {
				if x.Op.String() == "&" {
					if store[funcName][lVarName].typeVar[0] == '*' {
						store.addNewValue(funcName, lVarName, "valid", lVarPos)
					}
				}
			} else if node.Tok.String() == ":=" {
				if x.Op.String() == "&" {
					//add pointer type
					store.addNewVar(funcName, lVarName, "*?", "valid", lVarPos)
				}
			}
		case *ast.Ident:
			if node.Tok.String() == "=" {
				if store[funcName][lVarName].typeVar[0] == '*' {
					value := store.getLastValue(funcName, x.Name)
					store.addNewValue(funcName, lVarName, value, lVarPos)
				}
			} else if node.Tok.String() == ":=" {
				_, keyIsValid := store[funcName][rVarName]
				if keyIsValid {
					value := store.getLastValue(funcName, rVarName)
					typeVar := store.getVarType(funcName, rVarName)
					store.addNewVar(funcName, lVarName, typeVar, value, lVarPos)
				}
			}
			
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
			analizeAssignStmt(x, funcName, store)
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
