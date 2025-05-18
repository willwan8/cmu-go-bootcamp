/*
// Helpful pseudocode for this lesson (while loops)
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

/*
// Helpful pseudocode for this lesson (for loops)
AnotherFactorial(n)
  p ← 1
  for every integer i between 1 and n
    p ← p·i
  return p
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

	var count uint = 10
	for ; count >= 0; count-- {
		fmt.Println(count)
		// hidden count--
		// integer underflow, once we subtract 1 from 0, the unsigned integer will go to the largest possible unsigned integer as a negative uint doesn't exist
		// therefore, this is an infinite loop!
	}

	fmt.Println("Blast off!")
}

// AnotherFactorial takes as input an integer n and returns n!
func AnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to AnotherFactorial().")
	}
	p := 1

	// the "for" keyword is also used for for loops
	for i := 1; i <= n; i++ {
		// i := 1 is called the "initialization step"
		// i <= n is called the "condition"
		// i++ is called the "postcondition"

		p = p * i

		// there is a hidden line of code: i++
	}

	// this will not work as i is out of scope (its scope is only in the for loop)
	//fmt.Println("i is currently", i)

	return p
}

// SumEven takes as input an integer k and returns the sum of all even positive integers up to and (possibly) including k.
func SumEven(k int) int {
	if k < 0 {
		panic("Error: negative input given to SumEven().")
	}

	sum := 0

	// for every even integer j between 2 and k, add j to sum
	for j := 2; j <= k; j += 2 {
		// j must be even
		sum += j
	}

	return sum
}

func YetAnotherFactorial(n int) int {
	if n < 0 {
		panic("Error: negative input given to YetAnotherFactorial().")
	}

	p := 1

	for i := n; i >= 1; i-- {
		p *= i
	}

	return p
}

// AnotherSum takes an integer n as input and returns the sum of the first n positive integers, using a for loop
func AnotherSum(n int) int {
	if n < 0 {
		panic("Error: negative input given to AnotherSum().")
	}

	sum := 0 // will store the final answer

	for k := 1; k <= n; k++ {
		sum += k
	}

	return sum
}

func GaussSum(n int) int {
	return n * (n + 1) / 2
}

// mathematical fact: n! = n * (n-1)!
// when n = 1: 1! = 1 * 0!
// so 1 = 1 * 0! and therefore 1 = 0!

// Factorial takes as input an integer n and returns n! = n * (n - 1) * (n - 2) * ... * 2 * 1
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

	// i is still in scope (its scope is the function)
	fmt.Println("i is currently", i)

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