package main

import (
	"errors"
	"fmt"
	// "go/ast"
	// "go/parser"
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