package environment

import (
	"fmt"

	"github.com/5anthosh/chili/function"
	"github.com/shopspring/decimal"
)

//Type of variables
const (
	functionType = 1 + iota
	variableType
)

//Environment of evaluator
type Environment struct {
	symbolTable map[string]uint
	variables   map[string]interface{}
	functions   map[string]function.Function
}

//New Environment
func New() *Environment {
	return &Environment{
		symbolTable: make(map[string]uint),
		variables:   make(map[string]interface{}),
		functions:   make(map[string]function.Function),
	}
}

//SetFunction to Evaluator
func (e *Environment) SetFunction(function function.Function) error {
	if e.CheckSymbolTable(function.Name) {
		return fmt.Errorf("%s is already declared", function.Name)
	}
	e.symbolTableEntry(function.Name, functionType)
	e.functions[function.Name] = function
	return nil
}

//SetDefaultFunctions to environment
func (e *Environment) SetDefaultFunctions() error {
	err := e.SetFunction(function.AbsFunction)
	return err
}

//SetNumberVariable in the environment
func (e *Environment) SetNumberVariable(name string, value decimal.Decimal) error {
	return e.DeclareVariable(name, value)
}

//DeclareVariable in the environment
func (e *Environment) DeclareVariable(name string, value interface{}) error {
	if e.CheckSymbolTable(name) {
		return fmt.Errorf("%s is already declared", name)
	}
	e.symbolTableEntry(name, variableType)
	e.variables[name] = value
	return nil
}

//GetVariable from the environment
func (e *Environment) GetVariable(name string) (interface{}, bool) {
	value, ok := e.variables[name]
	return value, ok
}

//GetFunction from the environment
func (e *Environment) GetFunction(name string) (function.Function, bool) {
	value, ok := e.functions[name]
	return value, ok
}

func (e *Environment) symbolTableEntry(name string, varType uint) {
	e.symbolTable[name] = varType
}

//CheckSymbolTable if variable is registered in symbol table
func (e *Environment) CheckSymbolTable(name string) bool {
	_, ok := e.symbolTable[name]
	return ok
}

//IsVariable check if name is variable
func (e *Environment) IsVariable(name string) bool {
	varType, _ := e.symbolTable[name]
	return varType == variableType
}

//IsFunction check if name is variable
func (e *Environment) IsFunction(name string) bool {
	varType, _ := e.symbolTable[name]
	return varType == functionType
}
