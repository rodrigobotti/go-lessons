package math

import (
	"fmt"
	"reflect"
	"runtime"
)

// Calculate executes the operation and returns its results
func Calculate(operation func(int, int) int, x int, y int) int {
	opName := runtime.FuncForPC(reflect.ValueOf(operation).Pointer()).Name()
	fmt.Printf("Running %v with parameters (%d, %d)\n", opName, x, y)
	return operation(x, y)
}
