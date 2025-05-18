/*
// Helpful pseudocode for this lesson
Factorial(n)
  p ← 1
  i ← 1
  while i ≤ n
    p ← p · i
    i ← i + 1
  return p

EuclidGCD(a,b)
  while a ≠ b
    if a > b
      a ← a − b
    else
      b ← b − a
  return a 
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops!")

	fmt.Println(Factorial(5))

	fmt.Println(Factorial(0))

	// these factorials are so big that it results in integer overflow, so we get unexpected values!
	fmt.Println(Factorial(40))
	fmt.Println(Factorial(45))
	fmt.Println(Factorial(300))
}

// mathematical fact: n! = n * (n-1)!
// when n = 1: 1! = 1 * 0!
// so 1 = 1 * 0! and therefore 1 = 0!

//Factorial takes as input an integer n and returns n! = n * (n - 1) * (n - 2) * ... * 2 * 1
func Factorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to Factorial().") //Go's error message function
	}

	p := 1 // will store product
	i := 1 // serves as a counter

	// Go has no keyword "while" and uses "for" instead
	for i <= n {
		p = p * i
		i = i + 1
	}

	return p
}

// SumFirstNIntegers takes as input an integer n and returns the sum of the first n positive integers.
func SumFirstNIntegers(n int) int {
	if n < 0 {
		panic("Error: negative input given to SumFirstNIntegers()")
	}

	sum := 0 // will store the final answer
	i := 1 // represents current integer being added to sum

	for i <= n {
		sum += i // shorthand for sum = sum + i
		// also have sum -= i, sum *= i, sum /= i
		i++ // shorthand for i = i + 1
		// there is also i-- for i = i - 1
	}

	return sum
}