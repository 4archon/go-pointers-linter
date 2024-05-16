package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func tokenDir(scandir string, fset *token.FileSet) error {
	dir, err := os.ReadDir(scandir)
	if err != nil {
		fmt.Println("Read dir error")
		return errors.New("Read dir error")
	}
	for _, file := range dir {
		var fileName string = scandir + file.Name()
		stat, err := os.Stat(fileName)
		if err != nil {
			fmt.Println("Read stst error")
			return errors.New("Read stst error")
		}
		(*fset).AddFile(fileName, -1, int(stat.Size()))
	}
	return nil
}

type varProg struct {
	name string
	typeVar string
	value *int
	pos token.Pos
}

// func getLine(fset token.FileSet, pos token.Pos) int {
	
// }

type storage map[string]map[string][]varProg

func getType(node *ast.StarExpr) (string) {
	switch x := node.X.(type) {
	case *ast.Ident:
		return "*" + x.Name
	}
	return ""
}

func (store *storage) analizeDeclStmt(node *ast.GenDecl, funcName string) {
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
				(*store)[funcName][x.Names[0].Name] = append((*store)[funcName][x.Names[0].Name], s)
			}
		}
	}
}

func (store *storage) analizeFuncBody(node *ast.BlockStmt, funcName string) {
	(*store)[funcName] = make(map[string][]varProg)
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

func (store *storage) storeVar(node ast.Node) bool {
	switch x := node.(type) {
	case *ast.FuncDecl:
		(*store).analizeFuncBody(x.Body, x.Name.Name)
	}
	return true
}

func getLine(fset *token.FileSet, pos token.Pos) int {
	return fset.Position(pos).Line
}

func (store storage) printStore(fset *token.FileSet) {
	for _, i := range store {
		for _, j := range i {
			for _, k := range j {
				fmt.Println(k.name)
				fmt.Println(getLine(fset, k.pos))
				fmt.Println(k.typeVar)
				fmt.Println(k.value)
				fmt.Println("--------")
			}
		}
	}
}

// func ge

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
		s := make(storage)
		ast.Inspect(node["main"], s.storeVar)
		// fmt.Println(s["main"]["i2"][0].)
		s.printStore(fset)
	}

}
