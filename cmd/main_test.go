package main

import "testing"

var result bool

func BenchmarkSolutionA(b *testing.B) {
	filePathA := "../resources/exported-data-before.json"
	filePathB := "../resources/exported-data-after.json"

	for i := 0; i < 5; i++ {
		result = SolutionA(filePathA, filePathB)
	}
}
