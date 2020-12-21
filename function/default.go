package function

import (
	"math"
	"strings"

	"github.com/5anthosh/chili/evaluator/datatype"
	"github.com/shopspring/decimal"
)

var one = decimal.NewFromInt(1)
var three = decimal.NewFromInt(3)
var two = decimal.NewFromInt(2)

//E Euler constant
var E, _ = decimal.NewFromString("2.71828182845904523536028747135266249775724709369995957496696763")

//PI constant
var PI, _ = decimal.NewFromString("3.14159265358979323846264338327950288419716939937510582097494459")

func absImpl(args []interface{}) (interface{}, error) {
	arg := args[0]
	return arg.(decimal.Decimal).Abs(), nil
}

func cbrtImpl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	result, _ := root(arg.BigFloat(), 3).Float64()
	return decimal.NewFromFloat(result), nil
}

func ceilImpl(args []interface{}) (interface{}, error) {
	arg := args[0]
	return arg.(decimal.Decimal).Ceil(), nil
}

func floorImpl(args []interface{}) (interface{}, error) {
	arg := args[0]
	return arg.(decimal.Decimal).Floor(), nil
}

func expImpl(args []interface{}) (interface{}, error) {
	arg := args[0]
	return E.Pow(arg.(decimal.Decimal)), nil
}

// TODO implement natural log using decimal.Decimal
func lnImpl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	v, _ := arg.Float64()
	return decimal.NewFromFloat(math.Log(v)), nil
}

func log10Impl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	v, _ := arg.Float64()
	return decimal.NewFromFloat(math.Log10(v)), nil
}

func log2Impl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	v, _ := arg.Float64()
	return decimal.NewFromFloat(math.Log2(v)), nil
}

func maxImpl(args []interface{}) (interface{}, error) {
	max := args[0].(decimal.Decimal)
	for i := 1; i < len(args); i++ {
		v := args[i].(decimal.Decimal)
		if v.GreaterThan(max) {
			max = v
		}
	}
	return max, nil
}

func minImpl(args []interface{}) (interface{}, error) {
	min := args[0].(decimal.Decimal)
	for i := 1; i < len(args); i++ {
		v := args[i].(decimal.Decimal)
		if v.LessThan(min) {
			min = v
		}
	}
	return min, nil
}

func powImpl(args []interface{}) (interface{}, error) {
	base := args[0].(decimal.Decimal)
	power := args[1].(decimal.Decimal)
	return base.Pow(power), nil
}

func roundImpl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	v, _ := arg.Float64()
	return decimal.NewFromFloat(math.Round(v)), nil
}

func signImpl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	return decimal.NewFromInt(int64(arg.Sign())), nil
}

func sqrtImpl(args []interface{}) (interface{}, error) {
	arg := args[0].(decimal.Decimal)
	result, _ := root(arg.BigFloat(), 2).Float64()
	return decimal.NewFromFloat(result), nil
}

func toNumberImpl(args []interface{}) (interface{}, error) {
	arg := args[0].(string)
	return decimal.NewFromString(arg)
}

/**************************String function************************/

func concatImpl(args []interface{}) (interface{}, error) {
	var builder strings.Builder
	for _, arg := range args {
		_, err := builder.WriteString(arg.(string))
		if err != nil {
			return "", nil
		}
	}
	return builder.String(), nil
}

func containsImpl(args []interface{}) (interface{}, error) {
	str1 := args[0].(string)
	str2 := args[1].(string)
	return strings.Contains(str1, str2), nil
}

func joinImpl(args []interface{}) (interface{}, error) {
	delimiter := args[0].(string)
	var builder strings.Builder
	_, err := builder.WriteString(args[1].(string))
	if err != nil {
		return nil, err
	}
	for i := 2; i < len(args); i++ {
		_, err = builder.WriteString(delimiter)
		if err != nil {
			return nil, err
		}
		_, err = builder.WriteString(args[i].(string))
		if err != nil {
			return nil, err
		}
	}
	return builder.String(), nil
}

func lengthImpl(args []interface{}) (interface{}, error) {
	return len(args[0].(string)), nil
}

func replaceImpl(args []interface{}) (interface{}, error) {
	str1 := args[0].(string)
	str2 := args[1].(string)
	str3 := args[2].(string)
	oc := args[3].(decimal.Decimal)
	return strings.Replace(str1, str2, str3, int(oc.IntPart())), nil
}

func replaceAllImpl(args []interface{}) (interface{}, error) {
	str1 := args[0].(string)
	str2 := args[1].(string)
	str3 := args[2].(string)
	return strings.ReplaceAll(str1, str2, str3), nil
}

func sliceImpl(args []interface{}) (interface{}, error) {
	str := args[0].(string)
	start := args[1].(decimal.Decimal).IntPart()
	stop := args[2].(decimal.Decimal).IntPart()
	if start < 0 {
		start = 0
	}
	if stop > int64(len(str)) {
		stop = int64(len(str))
	}
	return str[start:stop], nil
}

// Functions
var (
	AbsFunction = Function{
		Name:          "abs",
		Arity:         1,
		FunctionImpl:  absImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the absolute value of a number.\n Returns a number.",
	}
	CbrtFunction = Function{
		Name:          "cbrt",
		Arity:         1,
		FunctionImpl:  cbrtImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the cube root of a number.\n Returns a number.",
	}
	CeilFunction = Function{
		Name:          "ceil",
		Arity:         1,
		FunctionImpl:  ceilImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the smallest integer greater than or equal to a number.\n Returns a number.",
	}
	ExpFunction = Function{
		Name:          "exp",
		Arity:         1,
		FunctionImpl:  expImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns E^x, where x is the argument, and E is Euler's constant (2.718…), the base of the natural logarithm.\n Returns a number.",
	}
	FloorFunction = Function{
		Name:          "floor",
		Arity:         1,
		FunctionImpl:  floorImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the largest integer less than or equal to a number.\n Returns a number.",
	}
	LnFunction = Function{
		Name:          "ln",
		Arity:         1,
		FunctionImpl:  lnImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the natural logarithm of a number.\n Returns a number.",
	}
	Log10Function = Function{
		Name:          "log10",
		Arity:         1,
		FunctionImpl:  log10Impl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the base 10 logarithm of a number.\n Returns a number.",
	}
	Log2Function = Function{
		Name:          "log2",
		Arity:         1,
		FunctionImpl:  log2Impl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the base 2 logarithm of a number.\n Returns a number.",
	}
	MaxFunction = Function{
		Name:          "max",
		Arity:         -1,
		MinArity:      1,
		MaxArity:      MaximumNumberOfParamsLimit,
		FunctionImpl:  maxImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the largest number from a list, where numbers are separated by commas.\n Returns a number",
	}
	MinFunction = Function{
		Name:          "min",
		Arity:         -1,
		MinArity:      1,
		MaxArity:      MaximumNumberOfParamsLimit,
		FunctionImpl:  minImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the smallest number from a list, where numbers are separated by commas.\n Returns a number.",
	}
	PowFunction = Function{
		Name:          "pow",
		Arity:         2,
		FunctionImpl:  powImpl,
		ParamsType:    []uint{datatype.NumberType, datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns base to the exponent power.\n Returns a number.",
	}
	RoundFunction = Function{
		Name:          "round",
		Arity:         1,
		FunctionImpl:  roundImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Rounds a number to the nearest integer.\n Returns a number.",
	}
	SignFunction = Function{
		Name:          "sign",
		Arity:         1,
		FunctionImpl:  signImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the sign of a number, indicating whether it's positive (1), negative (-1) or zero (0).\n Returns a number.",
	}
	SqrtFunction = Function{
		Name:          "sqrt",
		Arity:         1,
		FunctionImpl:  sqrtImpl,
		ParamsType:    []uint{datatype.NumberType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the square root of a number.\n Returns a number.",
	}
	TonumberFunction = Function{
		Name:          "toNumber",
		Arity:         1,
		FunctionImpl:  toNumberImpl,
		ParamsType:    []uint{datatype.StringType},
		ReturnType:    datatype.NumberType,
		Documentation: "Converts a text string to a number.\n Returns a number.",
	}

	ConcantFunction = Function{
		Name:          "concat",
		Arity:         -1,
		MinArity:      1,
		MaxArity:      MaximumNumberOfParamsLimit,
		FunctionImpl:  concatImpl,
		ParamsType:    []uint{datatype.StringType},
		ReturnType:    datatype.StringType,
		Documentation: "Concatenates (combines) text strings.\n Returns a text string.",
	}

	ContainsFunction = Function{
		Name:          "contains",
		Arity:         2,
		FunctionImpl:  containsImpl,
		ParamsType:    []uint{datatype.StringType, datatype.StringType},
		ReturnType:    datatype.BooleanType,
		Documentation: "Tests whether a text string contains another text string.\n Returns a boolean.",
	}

	JoinFunction = Function{
		Name:          "join",
		Arity:         -1,
		MinArity:      2,
		MaxArity:      MaximumNumberOfParamsLimit,
		FunctionImpl:  joinImpl,
		ParamsType:    []uint{datatype.StringType},
		ReturnType:    datatype.StringType,
		Documentation: "Combines text strings, with a specified delimiter.\n Returns a text string.",
	}

	LengthFunction = Function{
		Name:          "length",
		Arity:         1,
		FunctionImpl:  lengthImpl,
		ParamsType:    []uint{datatype.StringType},
		ReturnType:    datatype.NumberType,
		Documentation: "Returns the number of characters in a text string.\n Returns a number.",
	}

	ReplaceFunction = Function{
		Name:          "replace",
		Arity:         4,
		FunctionImpl:  replaceImpl,
		ParamsType:    []uint{datatype.StringType, datatype.StringType, datatype.StringType, datatype.NumberType},
		ReturnType:    datatype.StringType,
		Documentation: "Replaces the n match within a text string with a specified new text string.\n Returns a text string.",
	}
	ReplaceAllFunction = Function{
		Name:          "replaceAll",
		Arity:         3,
		FunctionImpl:  replaceAllImpl,
		ParamsType:    []uint{datatype.StringType, datatype.StringType, datatype.StringType},
		ReturnType:    datatype.StringType,
		Documentation: "Replaces all matches within a text string with a specified new text string.\n Returns a text string.",
	}

	SliceFunction = Function{
		Name:          "slice",
		Arity:         3,
		FunctionImpl:  sliceImpl,
		ParamsType:    []uint{datatype.StringType, datatype.NumberType, datatype.NumberType},
		ReturnType:    datatype.StringType,
		Documentation: "Extracts a substring from a text string, given a specified starting point and  end point.\n Returns a text string",
	}
	DefaultFunctions = []Function{
		AbsFunction,
		CbrtFunction,
		CeilFunction,
		ExpFunction,
		FloorFunction,
		LnFunction,
		Log2Function,
		Log10Function,
		MaxFunction,
		MinFunction,
		PowFunction,
		RoundFunction,
		SignFunction,
		SqrtFunction,
		TonumberFunction,
		ConcantFunction,
		ContainsFunction,
		JoinFunction,
		LengthFunction,
		ReplaceFunction,
		ReplaceAllFunction,
		SliceFunction,
	}
)
