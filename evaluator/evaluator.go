package evaluator

import (
	"fmt"

	"github.com/5anthosh/eval/parser/ast/expr"
	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

//Evaluator #
type Evaluator struct {
	AST expr.Expr
}

//Run the evaluator
func (eval *Evaluator) Run() (interface{}, error) {
	return eval.accept(eval.AST)
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

	switch binaryExpr.Operator.Type {
	case token.Plus:
		return left.(decimal.Decimal).Add(right.(decimal.Decimal)), nil
	case token.Minus:
		return left.(decimal.Decimal).Sub(right.(decimal.Decimal)), nil
	case token.Star:
		return left.(decimal.Decimal).Mul(right.(decimal.Decimal)), nil
	case token.CommonSlash:
		return left.(decimal.Decimal).Div(right.(decimal.Decimal)), nil
	}

	return nil, fmt.Errorf("Unexpected binary operator %s", token.TokenVsTokenLiteral[binaryExpr.Operator.Type])
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
	if unaryExpr.Operator.Type == token.Minus {
		return (right.(decimal.Decimal)).Neg(), nil
	}
	return right, nil
}

func (eval *Evaluator) accept(expr expr.Expr) (interface{}, error) {
	return expr.Accept(eval)
}
