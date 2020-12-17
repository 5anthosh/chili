package ast

import (
	"fmt"
	"strings"

	"github.com/5anthosh/eval/parser/ast/expr"
)

const (
	tab    uint = 2
	preFix      = "+"
)

func createPrefix(depth uint, expressionType string) string {
	return fmt.Sprintf("%s%s (%d)%s", preFix, strings.Repeat("-", int(depth*tab)), depth/tab, expressionType)
}

//Printer #
type Printer struct {
	depth uint
}

//Print ast structure
func (ac *Printer) Print(expression expr.Expr) (string, error) {
	value, err := ac.accept(expression)
	if err != nil {
		return "", err
	}
	return value.(string), nil
}

//VisitBinaryExpr #
func (ac *Printer) VisitBinaryExpr(binaryExpr *expr.Binary) (interface{}, error) {
	ac.depth += tab
	left, err := ac.accept(binaryExpr.Left)
	if err != nil {
		return nil, err
	}
	right, err := ac.accept(binaryExpr.Right)
	if err != nil {
		return nil, err
	}
	ac.depth -= tab
	return fmt.Sprintf("%s %s \n|\n%v%v", createPrefix(ac.depth, "BINARY"), binaryExpr.Operator.Type.String(), left, right), nil
}

//VisitGroupExpr #
func (ac *Printer) VisitGroupExpr(groupExpr *expr.Group) (interface{}, error) {
	ac.depth += tab
	expression, err := ac.accept(groupExpr.Expression)
	if err != nil {
		return nil, err
	}
	ac.depth -= tab
	return fmt.Sprintf("%s \n|\n%v", createPrefix(ac.depth, "GROUP"), expression), nil
}

//VisitLiteralExpr #
func (ac *Printer) VisitLiteralExpr(literalExpression *expr.Literal) (interface{}, error) {
	return fmt.Sprintf("%s %v\n|\n", createPrefix(ac.depth, "LITERAL"), literalExpression.Value), nil
}

//VisitUnaryExpr #
func (ac *Printer) VisitUnaryExpr(unaryExpr *expr.Unary) (interface{}, error) {
	ac.depth += tab
	expression, err := ac.accept(unaryExpr)
	if err != nil {
		return nil, err
	}
	ac.depth -= tab
	return fmt.Sprintf("%s %s \n|\n%v", createPrefix(ac.depth, "UNARY"), unaryExpr.Operator.Type.String(), expression), nil
}

//VisitVariableExpr #
func (ac *Printer) VisitVariableExpr(variableExpression *expr.Variable) (interface{}, error) {
	return fmt.Sprintf("%s %s\n|\n", createPrefix(ac.depth, "VARIABLE"), variableExpression.Name), nil
}

func (ac *Printer) accept(expression expr.Expr) (interface{}, error) {
	return expression.Accept(ac)
}
