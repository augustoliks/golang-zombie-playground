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

func returnEmpty(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b

	return
}

func changeGlobalVar() {
	globalLevelVar01 = 456
	fmt.Println("In Function: ", globalLevelVar01)
	fmt.Println("In Function &: ", &globalLevelVar01)
}

func changeInternVar(a string) {
	a = a + "NETO"
	fmt.Println("In Function: ", a)
	fmt.Println("In Function &: ", &a)
}

func changeVariablePointer(a *string) {
	*a = *a + "NETO"
	fmt.Println("In Function: ", *a)
	fmt.Println("In Function &: ", a)
}

func main() {
	fmt.Println("\n---------------------")
	fmt.Println("Simple Function")
	sum := addVariables(5, 5)
	fmt.Println(sum)

	fmt.Println("\n---------------------")
	fmt.Println("Multiple Returns")
	value01, value02 := multipleReturns("carlos", "neto", 5, 5)
	fmt.Println(value01, value02)

	fmt.Println("\n---------------------")
	fmt.Println("Return Empty")
	splitArg01, splitArg02 := returnEmpty(5, 10)
	fmt.Println("Soma", splitArg01, "Subtração", splitArg02)

	fmt.Println("\n---------------------")
	fmt.Println(">>> Internal Variable")
	internalVariable := "Carlos"
	fmt.Println("Before: ", internalVariable)
	fmt.Println("Before &: ", &internalVariable)
	changeInternVar(internalVariable)
	fmt.Println("After: ", internalVariable)
	fmt.Println("After &: ", &internalVariable)

	fmt.Println("\n---------------------")
	fmt.Println(">>> Pointer Function Modify Variables")
	variablePointer := "Hamilton"
	fmt.Println("Before: ", variablePointer)
	fmt.Println("Before &:", &variablePointer)
	changeVariablePointer(&variablePointer)
	fmt.Println("After: ", variablePointer)
	fmt.Println("After &: ", &variablePointer)

	fmt.Println("\n---------------------")
	fmt.Println(">>> Change Global Variable")
	fmt.Println("Before: ", globalLevelVar01)
	fmt.Println("Before &:", &globalLevelVar01)
	changeGlobalVar()
	fmt.Println("After: ", globalLevelVar01)
	fmt.Println("After &: ", &globalLevelVar01)
}
