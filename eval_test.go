package eval

import (
	"fmt"
	"testing"

	"github.com/5anthosh/eval/environment"
	"github.com/5anthosh/eval/evaluator"
	"github.com/5anthosh/eval/parser"
	"github.com/shopspring/decimal"
)

func TestEvalNumberFunction(t *testing.T) {
	source := "PI*R^2 + abs(45.345)"
	env := environment.New()
	env.SetDefaultFunctions()
	env.SetNumberVariable("PI", decimal.RequireFromString("3.14"))
	env.SetNumberVariable("R", decimal.RequireFromString("2"))
	_parser := parser.New(source)
	_ast, err := _parser.Parse()
	if err != nil {
		t.Error(err)
	}
	_evaluator := evaluator.New(env, _ast)
	value, err := _evaluator.Run()
	if err != nil {
		t.Error(err)
	}
	if fmt.Sprintf("%v", value) != "57.905" {
		t.Errorf("Evaluation = %s but want %s", fmt.Sprintf("%v", value), "57.905")
	}
}
