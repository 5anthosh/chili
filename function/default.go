package function

import (
	"fmt"

	"github.com/5anthosh/eval/evaluator/datatype"
	"github.com/shopspring/decimal"
)

func checkNumber(val interface{}) bool {
	switch val.(type) {
	case decimal.Decimal:
		return true
	}
	return false
}

//AbsImpl #
func absImpl(args []interface{}) (interface{}, error) {
	arg := args[0]
	if !checkNumber(arg) {
		return nil, fmt.Errorf("Abs() is expecting number arg but got %v", arg)
	}
	return arg.(decimal.Decimal).Abs(), nil
}

// Functions
var (
	AbsFunction = Function{
		Name:         "abs",
		Arity:        1,
		FunctionImpl: absImpl,
		ParamsType:   []uint{datatype.NumeberType},
	}
)
