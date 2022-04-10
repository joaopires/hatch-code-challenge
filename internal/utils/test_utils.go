package utils

import (
	"fmt"
	"reflect"
)

func GetTestFailMessage(target string, expected interface{}, result interface{}) string {
	return fmt.Sprintf("%s - expected: %v, result: %v", target, result, expected)
}

func AssertErrors(expected, result error) bool {
	expectedErrorType := reflect.TypeOf(expected)
	resultErrType := reflect.TypeOf(result)

	return expectedErrorType != resultErrType
}
