package datatype

import (
	"errors"

	"github.com/shopspring/decimal"
)

//datatypes
const (
	NumeberType = 1 + iota
	StringType
	BooleanType
	GenerictType
	UnSupportedType
)

//ErrUnknownDataype #
var ErrUnknownDataype = errors.New("unknown datatype")

var typeVsString = []string{
	"NUMBER", "STRING", "BOOLEAN", "GENERIC", "UNSUPPORTED",
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

//CheckNumber #
func CheckNumber(values ...interface{}) bool {
	for _, value := range values {
		if !Checkdatatype(value, NumeberType) {
			return false
		}
	}
	return true
}

//CheckString #
func CheckString(values ...interface{}) bool {
	for _, value := range values {
		if !Checkdatatype(value, StringType) {
			return false
		}
	}
	return true
}

//CheckBoolean #
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
		return NumeberType, true
	case string:
		return StringType, true
	case bool:
		return BooleanType, true
	}
	return UnSupportedType, false
}

//GetTypeString #
func GetTypeString(value interface{}) string {
	dtype, _ := GetType(value)
	return typeVsString[int(dtype)-1]
}
