package evaluator

import (
	"errors"
	"fmt"

	"github.com/5anthosh/chili/environment"
	"github.com/5anthosh/chili/evaluator/datatype"
	"github.com/5anthosh/chili/parser/ast/expr"
	"github.com/5anthosh/chili/parser/token"
	"github.com/shopspring/decimal"
)

//errors
var (
	ErrDivisionByZero = errors.New("decimal division by zero")
)

//Evaluator #
type Evaluator struct {
	Env *environment.Environment
}

//New Evaluator
func New(env *environment.Environment) *Evaluator {
	if env == nil {
		env = environment.New()
	}
	return &Evaluator{
		Env: env,
	}
}

//Run the evaluator
func (eval *Evaluator) Run(expression expr.Expr) (interface{}, error) {
	return eval.accept(expression)
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
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).Add(right.(decimal.Decimal)), nil
		}
		if datatype.CheckString(left) && datatype.CheckNumber(right) {
			return left.(string) + (right.(decimal.Decimal)).String(), nil
		}
		if datatype.CheckNumber(left) && datatype.CheckString(right) {
			return (left.(decimal.Decimal)).String() + right.(string), nil
		}
		return nil, generateUnsupportedOperationErr("+", left, right)
	case token.MinusType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).Sub(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("-", left, right)
	case token.StarType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).Mul(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("*", left, right)
	case token.CommonSlashType:
		if datatype.CheckNumber(left, right) {
			if decimal.Zero.Equals(right.(decimal.Decimal)) {
				return nil, ErrDivisionByZero
			}
			return left.(decimal.Decimal).Div(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("/", left, right)
	case token.CapType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).Pow(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("^", left, right)
	case token.ModType:
		if datatype.CheckNumber(left, right) {
			if decimal.Zero.Equals(right.(decimal.Decimal)) {
				return nil, ErrDivisionByZero
			}
			return left.(decimal.Decimal).Mod(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("%", left, right)
	case token.EqualType:
		return logicalOperation("==", left, right)
	case token.NotEqualType:
		value, err := logicalOperation("!=", left, right)
		if err != nil {
			return nil, err
		}
		return !value, nil
	case token.GreaterType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).GreaterThan(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr(">", left, right)
	case token.GreaterEqualType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).GreaterThanOrEqual(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr(">=", left, right)
	case token.LesserType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).LessThan(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("<", left, right)
	case token.LesserEqualType:
		if datatype.CheckNumber(left, right) {
			return left.(decimal.Decimal).LessThanOrEqual(right.(decimal.Decimal)), nil
		}
		return nil, generateUnsupportedOperationErr("<=", left, right)
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

	if unaryExpr.Operator.Type.Type() == token.NotType {
		return !truthFullness(right), nil
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
	ok = eval.Env.IsVariable(variableExpr.Name)
	if !ok {
		return nil, fmt.Errorf("%s is not variable", variableExpr.Name)
	}
	value, ok := eval.Env.GetVariable(variableExpr.Name)
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

	err := _function.VerifyArgs(args)
	if err != nil {
		return nil, err
	}
	return _function.FunctionImpl(args)
}

//VisitTernary #
func (eval *Evaluator) VisitTernary(ternaryExpr *expr.Ternary) (interface{}, error) {
	cond, err := eval.accept(ternaryExpr.Condition)
	if err != nil {
		return nil, err
	}
	ok := truthFullness(cond)
	if ok {
		return eval.accept(ternaryExpr.True)
	}
	return eval.accept(ternaryExpr.False)
}

//VisitLogicalExpr #
func (eval *Evaluator) VisitLogicalExpr(logicalExpr *expr.Logical) (interface{}, error) {
	left, err := eval.accept(logicalExpr.Left)
	if err != nil {
		return nil, err
	}
	ok := truthFullness(left)
	switch logicalExpr.Operator.Type.Type() {
	case token.AndType:
		if !ok {
			return false, nil
		}
		right, err := eval.accept(logicalExpr.Right)
		if err != nil {
			return nil, err
		}
		return truthFullness(right), nil
	case token.OrType:
		if ok {
			return true, nil
		}
		right, err := eval.accept(logicalExpr.Right)
		if err != nil {
			return nil, err
		}
		return truthFullness(right), nil
	}
	return nil, fmt.Errorf("Unexpected logical operator %s", logicalExpr.Operator.Type.String())

}
func (eval *Evaluator) accept(expr expr.Expr) (interface{}, error) {
	return expr.Accept(eval)
}

func generateUnsupportedOperationErr(op string, left interface{}, right interface{}) error {
	return fmt.Errorf("%s operation between (%s, %s) is not supported", op, datatype.GetTypeString(left), datatype.GetTypeString(right))
}

func truthFullness(value interface{}) bool {
	if value == nil {
		return false
	}
	if datatype.CheckNumber(value) && value.(decimal.Decimal).Equals(decimal.Zero) {
		return true
	}
	if datatype.CheckString(value) && len(value.(string)) == 0 {
		return false
	}
	if datatype.CheckBoolean(value) {
		return value.(bool)
	}
	return true
}

func logicalOperation(op string, left interface{}, right interface{}) (bool, error) {
	if datatype.CheckNumber(left, right) {
		return left.(decimal.Decimal).Equals(right.(decimal.Decimal)), nil
	}
	if datatype.CheckString(left, right) {
		return left.(string) == right.(string), nil
	}
	if datatype.CheckBoolean(left, right) {
		return left.(bool) == right.(bool), nil
	}
	return false, generateUnsupportedOperationErr(op, left, right)
}
