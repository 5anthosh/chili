package chili

import (
	"fmt"
	"testing"

	"github.com/5anthosh/chili/environment"
	"github.com/5anthosh/chili/evaluator"
	"github.com/5anthosh/chili/parser"
	"github.com/shopspring/decimal"
)

func TestEvalNumberFunction(t *testing.T) {
	source := "PI*R^2 + abs(45.345)"
	env := environment.New()
	env.SetDefaultFunctions()
	env.SetNumberVariable("PI", decimal.RequireFromString("3.1415926535897932385"))
	env.SetNumberVariable("R", decimal.RequireFromString("2"))
	chiliParser := parser.New(source)
	expression, err := chiliParser.Parse()
	if err != nil {
		t.Error(err)
	}
	chiliEvaluator := evaluator.New(env)
	value, err := chiliEvaluator.Run(expression)
	if err != nil {
		t.Error(err)
	}
	if fmt.Sprintf("%v", value) != "57.911370614359172954" {
		t.Errorf("Evaluation = %s but want %s", fmt.Sprintf("%v", value), "57.905")
	}
}
