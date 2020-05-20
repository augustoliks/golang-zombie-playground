package main

import (
	"fmt"
)

var globalLevelVar01 int
var globalLevelVar02 string

func addVariables(a, b int) int {
	return a + b
}

func multipleReturns(a, b string, c, d int) (string, int) {
	return a + b, c + d
}

func splitArgsFunction(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b

	return
}

func changeGlobalVar() {
	globalLevelVar01 = 123
}

func changeInternVar(a string) {
	a = a + "NETO"
	fmt.Println("In Function &: ", &a)
	fmt.Println("In Function: ", a)
}

func changeVariablePointer(a *string) {
	*a = *a + "NETO"
	fmt.Println("In Function &: ", a)
	fmt.Println("In Function: ", *a)
}

func main() {
	sum := addVariables(5, 5)
	fmt.Println(sum)

	value01, value02 := multipleReturns("carlos", "neto", 5, 5)
	fmt.Println(value01, value02)

	splitArg01, splitArg02 := splitArgsFunction(5, 10)
	fmt.Println("Soma", splitArg01, "Subtração", splitArg02)

	internalVariable := "Carlos"
	fmt.Println("Before: ", internalVariable)
	fmt.Println("Before &: ", &internalVariable)
	changeInternVar(internalVariable)
	fmt.Println("After: ", internalVariable)
	fmt.Println("After &: ", &internalVariable)

	variablePointer := "Hamilton"
	fmt.Println("Before: ", variablePointer)
	fmt.Println("Before &:", &variablePointer)
	changeVariablePointer(&variablePointer)
	fmt.Println("After: ", variablePointer)
	fmt.Println("After &: ", &variablePointer)

	fmt.Println(globalLevelVar01)
	changeGlobalVar()
	fmt.Println(globalLevelVar01)

}
