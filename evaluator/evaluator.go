package evaluator

import (
	"fmt"

	"github.com/5anthosh/eval/evaluator/function"
	"github.com/5anthosh/eval/parser/ast/expr"
	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

const (
	functionType = 1 + iota
	variableType
)

//Evaluator #
type Evaluator struct {
	strict      bool
	symbolTable map[string]uint
	variables   map[string]interface{}
	functions   map[string]function.Function
}

//New Evaluator
func New(strict bool) *Evaluator {
	return &Evaluator{
		strict:      strict,
		symbolTable: make(map[string]uint),
		variables:   make(map[string]interface{}),
		functions:   make(map[string]function.Function),
	}
}

//SetFunction #
func (eval *Evaluator) SetFunction(function function.Function) error {
	if eval.checkSymbolTable(function.Name, functionType) {
		return fmt.Errorf("%s() is already declared", function.Name)
	}
	eval.symbolTableEntry(function.Name, functionType)
	eval.functions[function.Name] = function
	return nil
}

//SetNumberVariable #
func (eval *Evaluator) SetNumberVariable(name string, value decimal.Decimal) error {
	if eval.checkSymbolTable(name, variableType) {
		return fmt.Errorf("%s is already declared", name)
	}
	eval.symbolTableEntry(name, variableType)
	eval.variables[name] = value
	return nil
}

func (eval *Evaluator) symbolTableEntry(name string, symbolType uint) {
	eval.symbolTable[name] = symbolType
}

func (eval *Evaluator) checkSymbolTable(name string, symbolType uint) bool {
	t, ok := eval.symbolTable[name]
	if !ok {
		return ok
	}

	if t == symbolType || eval.strict && t != symbolType {
		return true
	}
	return false
}

//Run the evaluator
func (eval *Evaluator) Run(AST expr.Expr) (interface{}, error) {
	return eval.accept(AST)
}

//VisitBinaryExpr #
func (eval *Evaluator) VisitBinaryExpr(binaryExpr *expr.Binary) (interface{}, error) {
	left, err := eval.accept(binaryExpr.Left)
	if err != nil {
		return nil, err
	}
	right, err := eval.accept(binaryExpr.Right)
	if err != nil {
		return nil, err
	}

	switch binaryExpr.Operator.Type.Type() {
	case token.PlusType:
		return left.(decimal.Decimal).Add(right.(decimal.Decimal)), nil
	case token.MinusType:
		return left.(decimal.Decimal).Sub(right.(decimal.Decimal)), nil
	case token.StarType:
		return left.(decimal.Decimal).Mul(right.(decimal.Decimal)), nil
	case token.CommonSlashType:
		return left.(decimal.Decimal).Div(right.(decimal.Decimal)), nil
	}

	return nil, fmt.Errorf("Unexpected binary operator %s", binaryExpr.Operator.Type.String())
}

//VisitGroupExpr #
func (eval *Evaluator) VisitGroupExpr(groupExpr *expr.Group) (interface{}, error) {
	return eval.accept(groupExpr.Expression)
}

//VisitLiteralExpr #
func (eval *Evaluator) VisitLiteralExpr(literalExpression *expr.Literal) (interface{}, error) {
	return literalExpression.Value, nil
}

//VisitUnaryExpr #
func (eval *Evaluator) VisitUnaryExpr(unaryExpr *expr.Unary) (interface{}, error) {
	right, err := eval.accept(unaryExpr.Right)
	if err != nil {
		return nil, err
	}
	if unaryExpr.Operator.Type.Type() == token.MinusType {
		return (right.(decimal.Decimal)).Neg(), nil
	}
	return right, nil
}

//VisitVariableExpr #
func (eval *Evaluator) VisitVariableExpr(variableExpr *expr.Variable) (interface{}, error) {
	value, ok := eval.variables[variableExpr.Name]
	if ok {
		return value, nil
	}
	return nil, fmt.Errorf("Unknown variable %s", variableExpr.Name)
}

func (eval *Evaluator) accept(expr expr.Expr) (interface{}, error) {
	return expr.Accept(eval)
}
