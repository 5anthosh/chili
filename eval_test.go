package chili

import (
	"fmt"
	"testing"

	"github.com/5anthosh/chili/environment"
	"github.com/5anthosh/chili/evaluator"
	"github.com/5anthosh/chili/parser"
)

func TestEvalNumberFunction(t *testing.T) {
	source := "PI*R^2 + abs(45.345)"
	env := environment.New()
	env.SetDefaultFunctions()
	env.SetDefaultVariables()
	env.SetIntVariable("R", 2)
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
	if fmt.Sprintf("%v", value) != "57.91137061435917295385057353311801153678867759750042328389977836" {
		t.Errorf("Evaluation = %s but want %s", fmt.Sprintf("%v", value), "57.905")
	}
}
