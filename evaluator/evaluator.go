package evaluator

import (
	"fmt"

	"github.com/5anthosh/eval/evaluator/function"
	"github.com/5anthosh/eval/parser/ast/expr"
	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

//Evaluator #
type Evaluator struct {
	symbolTable map[string]struct{}
	variables   map[string]interface{}
	functions   map[string]function.Function
}

//New Evaluator
func New() *Evaluator {
	return &Evaluator{
		symbolTable: make(map[string]struct{}),
		variables:   make(map[string]interface{}),
		functions:   make(map[string]function.Function),
	}
}

//SetFunction to Evaluator
func (eval *Evaluator) SetFunction(function function.Function) error {
	if eval.checkSymbolTable(function.Name) {
		return fmt.Errorf("%s is already declared", function.Name)
	}
	eval.symbolTableEntry(function.Name)
	eval.functions[function.Name] = function
	return nil
}

func (eval *Evaluator) defaultFunctions() {
	eval.SetFunction(function.AbsFunction)
}

//SetNumberVariable #
func (eval *Evaluator) SetNumberVariable(name string, value decimal.Decimal) error {
	if eval.checkSymbolTable(name) {
		return fmt.Errorf("%s is already declared", name)
	}
	eval.symbolTableEntry(name)
	eval.variables[name] = value
	return nil
}

func (eval *Evaluator) symbolTableEntry(name string) {
	eval.symbolTable[name] = struct{}{}
}

func (eval *Evaluator) checkSymbolTable(name string) bool {
	_, ok := eval.symbolTable[name]
	return ok
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
