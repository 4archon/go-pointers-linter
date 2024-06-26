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

func getLine(fset *token.FileSet, pos token.Pos) int {
	return fset.Position(pos).Line
}

func cmpPos(lPos token.Pos, rPos token.Pos, fset *token.FileSet) (bool ,error) {
	lPosition := fset.Position(lPos)
	rPosition := fset.Position(rPos)
	if lPosition.Filename == rPosition.Filename {
		lrow := lPosition.Line
		rrow := rPosition.Line
		if lrow < rrow {
			return false, nil
		} else {
			return true, nil
		}
	} else {
		return false, errors.New("Not the same files")
	}
}