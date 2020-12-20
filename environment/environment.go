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
	funcs := function.DefaultFunctions
	for _, chiliFunc := range funcs {
		err := e.SetFunction(chiliFunc)
		if err != nil {
			return err
		}
	}
	return nil
}

//SetDefaultVariables to environment
func (e *Environment) SetDefaultVariables() error {
	err := e.SetNumberVariable("PI", function.PI)
	if err != nil {
		return err
	}
	err = e.SetNumberVariable("E", function.E)
	return err
}

//SetNumberVariable in the environment
func (e *Environment) SetNumberVariable(name string, value decimal.Decimal) error {
	return e.DeclareVariable(name, value)
}

//SetFloatVariable  in the environment
func (e *Environment) SetFloatVariable(name string, value float64) error {
	return e.SetNumberVariable(name, decimal.NewFromFloat(value))
}

//SetFloat32Variable in the environment
func (e *Environment) SetFloat32Variable(name string, value float32) error {
	return e.SetNumberVariable(name, decimal.NewFromFloat32(value))
}

//SetIntVariable in the environment
func (e *Environment) SetIntVariable(name string, value int64) error {
	return e.SetNumberVariable(name, decimal.NewFromInt(value))
}

//SetInt32Variable in the environment
func (e *Environment) SetInt32Variable(name string, value int32) error {
	return e.SetNumberVariable(name, decimal.NewFromInt32(value))
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
