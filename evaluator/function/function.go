package function

//Function struct
type Function struct {
	Name         string
	Arity        uint
	FunctionImpl func(argument []interface{}) (interface{}, error)
}
