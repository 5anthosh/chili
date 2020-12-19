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
	ast   expr.Expr
}

//New Printer
func New(ast expr.Expr) *Printer {
	return &Printer{
		depth: 0,
		ast:   ast,
	}
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
	expression, err := ac.accept(unaryExpr.Right)
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

//VisitFunctionCall #
func (ac *Printer) VisitFunctionCall(functionCallExpr *expr.FunctionCall) (interface{}, error) {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("%s %s\n|\n", createPrefix(ac.depth, "FUNCTION"), functionCallExpr.Name))
	ac.depth += tab
	for _, arg := range functionCallExpr.Args {
		argStr, err := ac.accept(arg)
		if err != nil {
			return nil, err
		}
		builder.WriteString(fmt.Sprintf("%s", argStr))
	}
	ac.depth -= tab
	return builder.String(), nil

}

//VisitTernary #
func (ac *Printer) VisitTernary(ternaryExpr *expr.Ternary) (interface{}, error) {
	ac.depth += tab
	cond, err := ac.accept(ternaryExpr.Condition)
	if err != nil {
		return nil, err
	}
	trueExpr, err := ac.accept(ternaryExpr.True)
	if err != nil {
		return nil, err
	}
	falseExpr, err := ac.accept(ternaryExpr.False)
	if err != nil {
		return nil, err
	}
	ac.depth -= tab
	return fmt.Sprintf("%s \n|\n %v%v%v", createPrefix(ac.depth, "TERNARY"), cond, trueExpr, falseExpr), nil
}

//VisitLogicalExpr #
func (ac *Printer) VisitLogicalExpr(logicalExpr *expr.Logical) (interface{}, error) {
	ac.depth += tab
	left, err := ac.accept(logicalExpr.Left)
	if err != nil {
		return nil, err
	}
	right, err := ac.accept(logicalExpr.Right)
	if err != nil {
		return nil, err
	}
	ac.depth -= tab
	return fmt.Sprintf("%s %s \n|\n%v%v", createPrefix(ac.depth, "LOGICAL"), logicalExpr.Operator.Type.String(), left, right), nil
}

func (ac *Printer) accept(expression expr.Expr) (interface{}, error) {
	return expression.Accept(ac)
}

func (ac *Printer) String() (string, error) {
	v, err := ac.accept(ac.ast)
	if err != nil {
		return "", err
	}
	return v.(string), err
}
