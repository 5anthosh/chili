package evaluator

import (
	"testing"

	"github.com/5anthosh/eval/parser"
	"github.com/shopspring/decimal"
)

func TestEvaluatorNormalExpression(t *testing.T) {
	expression := "34534 + 345.34 - 222 / 43435 * 745.234"
	parser := parser.New(expression)
	expr, err := parser.Parse()
	if err != nil {
		t.Error(err)
	}
	eval := New()
	value, err := eval.Run(expr)
	if err != nil {
		t.Error(err)
	}
	println(value.(decimal.Decimal).String())
}
