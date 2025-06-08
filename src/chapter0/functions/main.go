package main

import (
	"fmt"
)

/* Note: 
	the order of functions doesn't matter
*/
func main() {
	fmt.Println("Functions!")

	x := 3 // x has type int
	n := SumTwoInts(x, x)
	fmt.Println(n)

	m := 17
	fmt.Println(AddOne(m)) // a copy of m is sent as the input parameter to AddOne
	fmt.Println(m) // what gets printed? (17)
}

// to make a function, we need to declare them, just like with variables 
// function signature: func name(input parameters) return-type
// within the parentheses, we need to declare any parameters we want to take in (variable name followed by type)
// specify the return type of the function (ex: int)

// AddOne takes an interger k as input and returns the value of k + 1.
// m only has scope within this function (changes will not apply outside the function)
// go uses pass by value (when you call a function, a copy of the variable for the input is passed to the function)
func AddOne(m int) int {
	m = m + 1
	return m
}

// SumTwoInts takes two as input two integers and returns their sum.
// a shorthand for input parameters of the same type is: a, b int (int is assumed for all previous parameters)
func SumTwoInts(a int, b int) int {
	return a + b
}

// DoubleAndDuplicate takes as input a float64.
// It returns two copies of that input variable.
func DoubleAndDuplicate(x float64) (float64, float64) {
	return 2.0*x, 2.0*x
}

// A function doesn't have to have input parameters
// Pi takes no inputs and returns the value of pi, the mathematical constant.
func Pi() float64 {
	return 3.14
}

// A function doesn't have to have a return-type
// PrintHi simply prints a message "Hi" to the console.
func PrintHi() {
	fmt.Println("Hi")
}