package evaluator

import (
	"fmt"

	"github.com/5anthosh/eval/environment"
	"github.com/5anthosh/eval/parser/ast/expr"
	"github.com/5anthosh/eval/parser/token"
	"github.com/shopspring/decimal"
)

//Evaluator #
type Evaluator struct {
	Env *environment.Environment
	AST expr.Expr
}

//New Evaluator
func New(env *environment.Environment, AST expr.Expr) *Evaluator {
	if env == nil {
		env = environment.New()
	}
	return &Evaluator{
		Env: env,
		AST: AST,
	}
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

	switch binaryExpr.Operator.Type.Type() {
	case token.PlusType:
		return left.(decimal.Decimal).Add(right.(decimal.Decimal)), nil
	case token.MinusType:
		return left.(decimal.Decimal).Sub(right.(decimal.Decimal)), nil
	case token.StarType:
		return left.(decimal.Decimal).Mul(right.(decimal.Decimal)), nil
	case token.CommonSlashType:
		return left.(decimal.Decimal).Div(right.(decimal.Decimal)), nil
	case token.CapType:
		return left.(decimal.Decimal).Pow(right.(decimal.Decimal)), nil
	case token.ModType:
		return left.(decimal.Decimal).Mod(right.(decimal.Decimal)), nil
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
	ok := eval.Env.CheckSymbolTable(variableExpr.Name)
	if !ok {
		return nil, fmt.Errorf("Unknown variable %s", variableExpr.Name)
	}
	ok = eval.Env.IsVarible(variableExpr.Name)
	if !ok {
		return nil, fmt.Errorf("%s is not variable", variableExpr.Name)
	}
	value, ok := eval.Env.GetVarible(variableExpr.Name)
	if ok {
		return value, nil
	}
	return nil, fmt.Errorf("Unknown variable %s", variableExpr.Name)
}

//VisitFunctionCall #
func (eval *Evaluator) VisitFunctionCall(functionCall *expr.FunctionCall) (interface{}, error) {
	ok := eval.Env.CheckSymbolTable(functionCall.Name)
	if !ok {
		return nil, fmt.Errorf("Unknown function %s()", functionCall.Name)
	}
	ok = eval.Env.IsFunction(functionCall.Name)
	if !ok {
		return nil, fmt.Errorf("%s is not function", functionCall.Name)
	}
	_function, _ := eval.Env.GetFunction(functionCall.Name)

	var args []interface{}
	for _, arg := range functionCall.Args {
		value, err := eval.accept(arg)
		if err != nil {
			return nil, err
		}
		args = append(args, value)
	}
	ok = _function.CheckNumberOfArgs(args)
	if !ok {
		return nil, fmt.Errorf("function %s() expecting %d but got %d", functionCall.Name, _function.Arity, len(args))
	}
	ok = _function.CheckTypeOfArgs(args)
	if !ok {
		return nil, fmt.Errorf("function %s() got wrong data type params", functionCall.Name)
	}
	return _function.FunctionImpl(args)
}

func (eval *Evaluator) accept(expr expr.Expr) (interface{}, error) {
	return expr.Accept(eval)
}
