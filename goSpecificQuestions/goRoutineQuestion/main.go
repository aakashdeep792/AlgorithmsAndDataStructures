package main

import (
	"algoDS/goSpecificQuestions/goRoutineQuestion/code"
	"fmt"
)

func main() {
	// a := code.CheckPrime(1000) //code.PrintNumber4(15)
	// a := code.WordCountInFile("./goSpecificQuestions/goRoutineQuestion/code/calculateSum.go")
		a := code.ReadCSV("./code/customers-100.csv")
	fmt.Println(a)
}
