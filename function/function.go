package function

import (
	"fmt"

	"github.com/5anthosh/chili/evaluator/datatype"
)

//MaximumNumberOfParamsLimit for the function
const MaximumNumberOfParamsLimit = 255

//Function struct
type Function struct {
	Name               string
	Arity              int
	MinArity           uint
	MaxArity           uint
	FunctionImpl       func(args []interface{}) (interface{}, error)
	ParamsType         []uint
	VerifyArgs         func(arguments []interface{}) error
	ReturnType         uint
	Documentation      string
	ArgsDocumentation  string
	ExampleDocumention string
}

//CheckNumberOfArgs in the function
func (f *Function) CheckNumberOfArgs(arguments []interface{}) error {
	if f.Arity == -1 {
		if len(arguments) < int(f.MinArity) {
			return fmt.Errorf("%s() expecting minimum %d arguments", f.Name, f.MinArity)
		}
		if len(arguments) > int(f.MaxArity) {
			return fmt.Errorf("%s() expecting maximum %d arguments", f.Name, f.MaxArity)
		}
		return nil
	}
	if f.Arity != len(arguments) {
		return fmt.Errorf("%s() expecting %d arguments but got %d", f.Name, f.Arity, len(arguments))
	}
	return nil
}

//CheckTypeOfArgs in the function
func (f *Function) CheckTypeOfArgs(arguments []interface{}) bool {
	if f.Arity == -1 {
		dtype := f.ParamsType[0]
		for _, arg := range arguments {
			if !datatype.Checkdatatype(arg, dtype) {
				return false
			}
		}
	}

	paramsLen := len(f.ParamsType)
	argsLen := len(arguments)
	for i := 0; i < paramsLen && i < argsLen; i++ {
		if !datatype.Checkdatatype(arguments[i], f.ParamsType[i]) {
			return false
		}
	}
	return true
}
