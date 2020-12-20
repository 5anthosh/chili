package datatype

import (
	"errors"

	"github.com/shopspring/decimal"
)

//datatypes
const (
	NumberType = 1 + iota
	StringType
	BooleanType
	GenerictType
	NoneType
	UnSupportedType
)

//ErrUnknownDataype #
var ErrUnknownDataype = errors.New("unknown datatype")

var typeVsString = []string{
	"NUMBER", "STRING", "BOOLEAN", "GENERIC", "NONE", "UNSUPPORTED",
}

//Checkdatatype of value is correct
func Checkdatatype(value interface{}, datatype uint) bool {
	if datatype == GenerictType {
		return true
	}
	dtype, _ := GetType(value)
	return dtype == datatype
}

//IsSupported check if datatype is supported
func IsSupported(value interface{}) bool {
	_, ok := GetType(value)
	return ok
}

//CheckNumber checks whether values are number type
func CheckNumber(values ...interface{}) bool {
	for _, value := range values {
		if !Checkdatatype(value, NumberType) {
			return false
		}
	}
	return true
}

//CheckString checks whether values are string type
func CheckString(values ...interface{}) bool {
	for _, value := range values {
		if !Checkdatatype(value, StringType) {
			return false
		}
	}
	return true
}

//CheckBoolean checks whether values are boolean type
func CheckBoolean(values ...interface{}) bool {
	for _, value := range values {
		if !Checkdatatype(value, BooleanType) {
			return false
		}
	}
	return true
}

//GetType of value
func GetType(value interface{}) (uint, bool) {
	switch value.(type) {
	case decimal.Decimal:
		return NumberType, true
	case string:
		return StringType, true
	case bool:
		return BooleanType, true
	case nil:
		return NoneType, true
	}
	return UnSupportedType, false
}

//GetTypeString #
func GetTypeString(value interface{}) string {
	dtype, _ := GetType(value)
	return typeVsString[int(dtype)-1]
}
