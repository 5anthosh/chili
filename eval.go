package eval

import (
	"github.com/5anthosh/eval/environment"
	"github.com/5anthosh/eval/evaluator"
	"github.com/5anthosh/eval/parser"
)

//Run the expression
func Run(expression string) (interface{}, error) {
	env := environment.New()
	env.SetDefaultFunctions()
	_parser := parser.New(expression)
	ast, err := _parser.Parse()
	if err != nil {
		return nil, err
	}
	_evaluator := evaluator.New(env, ast)
	return _evaluator.Run()
}
