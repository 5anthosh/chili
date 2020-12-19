package chili

import (
	"github.com/5anthosh/chili/environment"
	"github.com/5anthosh/chili/evaluator"
	"github.com/5anthosh/chili/evaluator/datatype"
	"github.com/5anthosh/chili/parser"
	"github.com/shopspring/decimal"
)

//Eval the expression
func Eval(expression string, data map[string]interface{}) (interface{}, error) {
	env := environment.New()
	env.SetDefaultFunctions()
	putDataToEnv(env, data)
	_parser := parser.New(expression)
	ast, err := _parser.Parse()
	if err != nil {
		return nil, err
	}
	_evaluator := evaluator.New(env)
	return _evaluator.Run(ast)
}

func putDataToEnv(env *environment.Environment, data map[string]interface{}) error {
	for k, v := range data {
		switch v.(type) {
		case int64:
			env.DeclareVariable(k, decimal.NewFromInt(v.(int64)))
		case int32:
			env.DeclareVariable(k, decimal.NewFromInt32(v.(int32)))
		case float64:
			env.DeclareVariable(k, decimal.NewFromFloat(v.(float64)))
		case float32:
			env.DeclareVariable(k, decimal.NewFromFloat32(v.(float32)))
		case string, bool:
			env.DeclareVariable(k, v)
		default:
			return datatype.ErrUnknownDataype
		}
	}
	return nil
}
