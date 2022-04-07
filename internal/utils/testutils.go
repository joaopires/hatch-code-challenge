package utils

import "fmt"

func GetTestFailMessage(target string, expected interface{}, result interface{}) string {
	return fmt.Sprintf("%s - expected: %v, result: %v", target, result, expected)
}
