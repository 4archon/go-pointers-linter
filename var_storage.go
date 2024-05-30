package main

import (
	// "errors"
	"fmt"
	// "go/ast"
	// "go/parser"
	"go/token"
	// "os"
)

type varValue struct {
	value string
	pos token.Pos
}

type varStore struct {
	name string
	typeVar string
	last int
	listVar []varValue
}

type storage map[string]map[string]varStore

func (store *storage) init() {
	*store = make(storage)
}

func (store storage) init2lvl(str string) {
	store[str] = make(map[string]varStore)
}

func (store storage) getLastValue(funcName string, name string) string {
	res := store[funcName][name].listVar[store[funcName][name].last]
	return res.value
}

func (store storage) getVarType(funcName string, name string) string {
	return store[funcName][name].typeVar
}

func (store storage) getValueAndPosVarByIndex(funcName string, name string, index int) (string, token.Pos) {
	value := store[funcName][name].listVar[index].value
	pos := store[funcName][name].listVar[index].pos
	return value, pos
}

func (store storage) addNewValue(funcName string, name string, value string, pos token.Pos) {
	var s varStore
	s.name = store[funcName][name].name
	s.typeVar = store[funcName][name].typeVar
	s.last = store[funcName][name].last + 1
	s.listVar = store[funcName][name].listVar
	
	var v varValue
	v.value = value
	v.pos = pos

	s.listVar = append(s.listVar, v)
	
	store[funcName][name] = s
}

func (store storage) addNewVar(funcName string, name string, typeVar string, value string, pos token.Pos) {
	var s varStore
	s.name = name
	s.typeVar = typeVar
	s.last = -1
	s.listVar = nil

	store[funcName][s.name] = s

	store.addNewValue(funcName, name, value, pos)
}

func (store storage) getLastFromDeref(funcName string, deref derefPointer, fset *token.FileSet) int {
	varChanges := store[funcName][deref.varName].listVar
	for j, i := range(varChanges) {
		cmp, valid := cmpPos(i.pos, deref.pos, fset)
		if valid == nil {
			if cmp {
				res := j -1
				if res < 0 {
					fmt.Println("error. deref before decl")
				}
				return res
			}
		}
	}
	res := store[funcName][deref.varName].last
	return res
}

func (store storage) printStore(fset *token.FileSet) {
	for funcName, i := range store {
		fmt.Println(funcName)
		fmt.Println("----------------------")
		for _, j := range i {
			for _, k := range j.listVar {
				fmt.Println(j.name)
				fmt.Println(getLine(fset, k.pos))
				fmt.Println(j.typeVar)
				fmt.Println(k.value)
				fmt.Println("--------")
			}
		}
	}
}