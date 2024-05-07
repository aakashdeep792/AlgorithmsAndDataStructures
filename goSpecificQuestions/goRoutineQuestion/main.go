package main

import (
	"algoDS/goSpecificQuestions/goRoutineQuestion/code"
	"fmt"
)

func main() {
	// a := code.CheckPrime(1000) //code.PrintNumber4(15)
	a := code.WordCountInFile("./goSpecificQuestions/goRoutineQuestion/code/calculateSum.go")
	fmt.Println(a)
}
