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

// func (store storage) addFromAnother(nameTo string, nameFrom string, funcName string) {
// 	store[funcName]
// }

// func (store storage) addExist ()

func (store storage) printStore(fset *token.FileSet) {
	for _, i := range store {
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