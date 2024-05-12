package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"errors"
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

type list []string

func (l *list) analize(node ast.Node) bool {
	var name string;
	switch x := node.(type) {
	case *ast.FuncDecl:
		name = x.Name.Name
	}
	if name != "" {
		*l = append(*l, name)
	}
	return true
}

type varProg struct {
	name string
	typeVar string
	value string
	pos string
}

type storage map[string]map[string][]varProg

func (store *storage) storeVar(node *ast.Node) bool {
	
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
		var l list
		ast.Inspect(node["main"], l.analize)
		fmt.Println(l)
	}

}
