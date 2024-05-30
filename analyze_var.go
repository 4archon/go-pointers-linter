package main

import (
	// "errors"
	// "fmt"
	"go/ast"
	// "go/parser"
	"go/token"
	// "os"
)

func getType(node *ast.StarExpr) (string) {
	switch x := node.X.(type) {
	case *ast.Ident:
		return "*" + x.Name
	}
	return ""
}
//need to add nil
func analyzeSpec(node *ast.Ident, expr *ast.StarExpr, funcName string, store storage, val string) {
	name := node.Name
	typeVar := getType(expr)
	value := val
	pos := node.NamePos

	store.addNewVar(funcName, name, typeVar, value, pos)
}

func analyzeSpecAndValues(node *ast.Ident, expr ast.Expr, typeExpr *ast.StarExpr, funcName string, store storage) {
	var value string
	switch x := expr.(type) {
	case *ast.UnaryExpr:
		if x.Op.String() == "&" {
			value = "valid"
		}
	case *ast.Ident:
		if x.Name == "nil" {
			value = "nil"
		} else {
			value = store.getLastValue(funcName, x.Name)
		}
	case *ast.CallExpr:
		value = "func"
	}
	analyzeSpec(node, typeExpr, funcName, store, value)
}

func getFuncNameCallExpr(node *ast.CallExpr) string {
	switch x := node.Fun.(type) {
	case *ast.Ident:
		return x.Name
	}
	return ""
}

func analyzeDeclStmt(node *ast.GenDecl, funcName string, store storage, ret funcReturns) {
	for _, i := range node.Specs {
		switch x := i.(type) {
		case *ast.ValueSpec:
			if x.Values == nil {
				switch y := x.Type.(type) {
				case *ast.StarExpr:
					for _, j := range x.Names {
						analyzeSpec(j, y, funcName, store, "nil")
					}
				}
			} else {
				switch y := x.Type.(type) {
				case *ast.StarExpr:
					if len(x.Names) == len(x.Values) {
						for z, j := range x.Names {
							analyzeSpecAndValues(j, x.Values[z], y, funcName, store)
						}
					} else {
						if len(x.Values) == 1 {
							switch w := x.Values[0].(type) {
							case *ast.CallExpr:
								name := getFuncNameCallExpr(w)
								if len(ret[name]) == len(x.Names) {
									for _, j := range x.Names {
										analyzeSpec(j, y, funcName, store, "func")
									}
								}
							}
						}
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

func analyzeAssignStmt(node *ast.AssignStmt, funcName string, store storage, ret funcReturns) {
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
					if x.Name == "nil" {
						store.addNewValue(funcName, lVarName, "nil", lVarPos)
					} else {
						value := store.getLastValue(funcName, x.Name)
						store.addNewValue(funcName, lVarName, value, lVarPos)
					}
				}
			} else if node.Tok.String() == ":=" {
				_, keyIsValid := store[funcName][rVarName]
				if keyIsValid {
					value := store.getLastValue(funcName, rVarName)
					typeVar := store.getVarType(funcName, rVarName)
					store.addNewVar(funcName, lVarName, typeVar, value, lVarPos)
				}
			}
		case *ast.CallExpr:
			if len(node.Lhs) != len(node.Rhs) && len(node.Rhs) == 1 {
				nameFuncCalled := getFuncNameCallExpr(x)
				if len(ret[nameFuncCalled]) == len(node.Lhs) {
					if node.Tok.String() == "=" {
						for _, z := range node.Lhs {
							name := getIdentName(z)
							pos := getIdentPos(z)
							store.addNewValue(funcName, name, "func", pos)
						}
					} else if node.Tok.String() == ":=" {
						var typesFunc []string = ret[nameFuncCalled]
						for t, z := range node.Lhs {
							typeVar := typesFunc[t]
							if typeVar[0] == '*' {
								name := getIdentName(z)
								pos := getIdentPos(z)
								store.addNewVar(funcName, name, typeVar, "func", pos)
							}
						}
					}
				}
			} else {
				if node.Tok.String() == "=" {
					store.addNewValue(funcName, lVarName, "func", lVarPos)
				} else if node.Tok.String() == ":=" {
					nameFuncCalled := getFuncNameCallExpr(x)
					typeVar := ret[nameFuncCalled][0]
					store.addNewVar(funcName, lVarName, typeVar, "func", lVarPos)
				}
			}
		}
	}
}

func analyzeFuncBody(node *ast.BlockStmt, funcName string, store storage, ret funcReturns) {
	store.init2lvl(funcName)
	for _, i := range node.List {
		switch x := i.(type) {
		case *ast.DeclStmt:
			switch y := x.Decl.(type) {
			case *ast.GenDecl:
				analyzeDeclStmt(y, funcName, store, ret)
			}
		case *ast.AssignStmt:
			analyzeAssignStmt(x, funcName, store, ret)
		}
	}
}

func analyzeFunc(funcStore funcStorage, store storage, ret funcReturns) {
	for _, i := range(funcStore) {
		analyzeFuncBody(i.Body, i.Name.Name, store, ret)
	}
}
