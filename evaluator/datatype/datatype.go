package datatype

import "github.com/shopspring/decimal"

//datatypes
const (
	NumeberType = 1 + iota
	StringType
	BooleanType
	GenerictType
)

//Checkdatatype of value is correct
func Checkdatatype(value interface{}, datatype uint) bool {
	if datatype == GenerictType {
		return true
	}
	switch value.(type) {
	case decimal.Decimal:
		return datatype == NumeberType
	case string:
		return datatype == StringType
	case bool:
		return datatype == BooleanType
	}
	return false
}

//IsSupported check if datatype is supported
func IsSupported(value interface{}) bool {
	switch value.(type) {
	case decimal.Decimal:
		return true
	case string:
		return true
	case bool:
		return false
	}
	return false
}
