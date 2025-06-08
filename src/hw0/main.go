package main

import (
	"fmt"
	"math"
)

func main() {
	// HW 2.1
	fmt.Println("Homework problem 2.1: Permutation Statistic:")
	fmt.Println(Permutation(12, 1), Permutation(6, 6), Permutation(10, 4), Permutation(10, 0))
	fmt.Println("Homework problem 2.1: Combination Statistic:")
	fmt.Println(Combination(10, 0), Combination(10, 10), Combination(12, 1), Combination(8, 7), Combination(10, 4), Combination(10, 6))

	// HW 2.2
	fmt.Println("Homework problem 2.2: Power:")
	fmt.Println(Power(2, 3), Power(1, 100), Power(7856, 0), Power(-1, 3), Power(7000, 1), Power(5, 4), Power(10, 6))
	fmt.Println("Homework problem 2.2: SumProperDivisors:")
	fmt.Println(SumProperDivisors(6), SumProperDivisors(40), SumProperDivisors(2003), SumProperDivisors(1))

	// HW 2.3
	fmt.Println("Homework problem 2.3: FibonacciArray:")
	fmt.Println(FibonacciArray(0), FibonacciArray(1), FibonacciArray(30))
	/* fmt.Println("Homework problem 2.3: DividesAll")
	no checks for this one */
	/* fmt.Println("Homework problem 2.3: MaxIntegerArray")
	no checks for this one */
	fmt.Println("Homework problem 2.3: MaxIntegers:")
	fmt.Println(MaxIntegers(-47), MaxIntegers(1, 2, 3, 4, 5, 6, 7), MaxIntegers(7, 6, 5, 4, 3, 2, 1), MaxIntegers(-100, -100, -1, -100, -100, -100))
	fmt.Println("Homework problem 2.3: SumIntegers:")
	fmt.Println(SumIntegers(42), SumIntegers(0, 1, -1), SumIntegers(1, 2, 3, 4, 5, 6), SumIntegers(-1, -1, -1, -1, -1, -1, -1, -1, -1, -1))
	/* fmt.Println("Homework problem 2.3: GCDArray:")
	no checks for this one */

	// HW 2.4
	fmt.Println("Homework problem 2.4: IsPerfect:")
	fmt.Println(IsPerfect(1), IsPerfect(2), IsPerfect(6), IsPerfect(10), IsPerfect(14), IsPerfect(28), IsPerfect(300), IsPerfect(496), IsPerfect(8000), IsPerfect(8128))
	fmt.Println("Homework problem 2.4: NextPerfectNumber:")
	fmt.Println(NextPerfectNumber(0), NextPerfectNumber(5), NextPerfectNumber(6), NextPerfectNumber(27), NextPerfectNumber(28), NextPerfectNumber(495), NextPerfectNumber(496), NextPerfectNumber(8128))
	fmt.Println("Homework problem 2.4: ListMersennePrimes:")
	fmt.Println(ListMersennePrimes(2), ListMersennePrimes(6), ListMersennePrimes(7), ListMersennePrimes(60))

	// HW 2.5
	/* fmt.Println("Homework problem 2.5: NextTwinPrimes")
	no checks for this one */
}

// Homework 2.1
// Permutation finds the permutation statistic, the number of times we can order k objects from n objects
func Permutation(n, k int) int {
	return (Factorial(n)/Factorial(n-k)) // permutation statistic = n!/((n-k)!)
}

// Factorial finds n!, where n! is n*(n-1)*(n-2)*...*(1)
func Factorial(n int) int  {
	// 0! equals 1 since n!=n*(n-1)!, so 1! = 1*0!, so 0!=1
	if n == 0 {
		return 1
	}
	// multiplies n by every number before it (n*(n-1)*(n-2)*...*(1))
	for i := n-1; i > 0; i-- {
		n *= i
	}

	return n
}

// Combination finds the combination statistic, the number of times we can select k from n objects while disregarding the arrangement of the objects
func Combination (n, k int) int {
	return (Factorial(n)/(Factorial(n-k)*Factorial(k))) // combination statistic = n!/((n-k)! * k!)
}

// Homework 2.2
// Power finds and returns a raised to the power of b
func Power(a, b int) int {
	result := 1

	for i := 0; i < b; i++ {
		result *= a
	}

	return result
}

// SumProperDivisors finds the sum all "proper" divisors of n (all divisors less than n)
func SumProperDivisors(n int) int {
	var sum int

	// iterates through all values from 1 to n and checks if it is a divisor of n
	for i := 1; i < n; i++ {
		if n % i == 0 { // if i is a divisor of n, add it to the sum
			sum += i
		}
	}

	return sum
}

// Homework 2.3
// FibonacciArray takes in n and returns an array of length n+1 whose k-th element is the k-th Fibnacci number
func FibonacciArray(n int) []int {
	fibonacci := make([]int, n+1)

	for i := 0; i <= n; i++ {
		if i <= 1 { // for the first two terms of fibonacci sequence 
			fibonacci[i] = 1
			continue // skips the rest of the code in for block
		}
		fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]
	}

	return fibonacci
}

// DividesAll takes an array of integers a and an integer d and returns true if d is a divisor of every integer in a, and false otherwise
func DividesAll(a []int, d int) bool {
	if d == 0 {
		return false
	}
	for i := 0; i < len(a); i++ { // checks through every integer in a
		if a[i] % d != 0 { // if d isn't a divisor of an integer in a, then we know that we should return false
			return false
		}
	}
	
	return true // all checks must have been passed, so d must be a divisor of all terms in a
}

// MaxIntegerArray takes an array of integers and returns the maximum value of all integers in the array
func MaxIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: list of length 0 given to MaxIntegerArray")
	}

	max := list[0]

	for i := 1; i < len(list); i++ {
		if list[i] > max {
			max = list[i]
		}
	}

	return max
}

// MaxIntegers takes an arbitrary number of integers and returns the max value of all the integers
func MaxIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: 0 integers given to MaxIntegers")
	}
	return MaxIntegerArray(numbers)
}

// SumIntegers takes an arbitrary colection of integers and returns the sum of all integers
func SumIntegers(numbers ...int) int {
	var sum int
	for _, val := range numbers { // equivalent to taking the value of each integer at each index in numbers
		sum += val
	}

	return sum
}

// GCDArray takes an array a and returns the greatest common divisor of all the integers in the array
func GCDArray(a []int) int {
	max := MaxIntegerArray(a)
	
	var d int
	for i := 1; i <= max; i++ { // goes through all possible values from 1 to the maximum value in the array
		if DividesAll(a, i) { // if i is a divisor of all integers, then we proceed
			d = i;
		}
	}

	return d
}

// Homework 2.4
// IsPerfect takes an integer n and returns if n is perfect (an integer that is equal to the sum of its proper divisors) or not
func IsPerfect(n int) bool {
	return n == SumProperDivisors(n) // if n is equal to the sum of its proper divisors, then it is perfect and returns true, else it returns false
}

// NextPerfectNumber takes an integer n and returns the smallest perfect number larger than n
func NextPerfectNumber(n int) int {
	var perfect int
	i := 1

	for IsPerfect(perfect) == false || perfect <= n { // will continue looping through possible numbers until we find a perfect number greater than n
		perfect = Power(2, i)*(Power(2, i+1)-1) // increases perfect according to our pattern (2^x * (2^(x+1) - 1))
		i++
	}

	return perfect
}

// ListMersennePrimes takes an integer n and returns an array of all primes of the form 2^p - 1, where p is a positive integer less than or equal to n
func ListMersennePrimes(n int) []int {
	var mersennePrimes []int

	for i := 2; i <= n; i++ { // i starts at 2 since 2^1 - 1 would be 1, and we want to exclude 1 as a prime for this problem
		if IsPrime(Power(2, i)-1) {
			mersennePrimes = append(mersennePrimes, (Power(2, i)-1))
		}
	}

	return mersennePrimes
}

// subroutine for ListMersennePrimes
// IsPrime takes as input an integer p and returns true if p is prime and false otherwise.
func IsPrime(p int) bool {
	for k := 2; float64(k) <= math.Sqrt(float64(p)); k++ {
		// is k a divisor of p? If so, return false
		if p % k == 0 {
			return false // composite
		}
	}
	// surviving all of these chekcs mens that p is prime
	return true
}

// Homework 2.5
// NextTwinPrimes takse an integer nand returns the smallest pair of twin primes (a pair of prime numbers only 2 apart) that are both larger than n
func NextTwinPrimes(n int) (int, int) {
	primeOne := n+1
	primeTwo := primeOne+2

	if primeOne == 1 {
		primeOne++
		primeTwo = primeOne+2
	}

	for IsPrime(primeOne) == false || IsPrime(primeTwo) == false { // both primeOne and primeTwo must be prime to exit this loop
		primeOne++
		primeTwo = primeOne+2 // increments so that primeTwo is always in a position to be a twin prime if possible
	}

	return primeOne, primeTwo
}