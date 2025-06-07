package main

import "fmt"

func main() {
	fmt.Println("Arrays (and slices).")

	var list [6]int

	// arrays are zero indexed
	list[0] = -8

	i := 3
	list[2*i-4] = 17 // index of 2

	list[len(list)-1] = 43

	// these commands lead to compiler errors (indexes are out of bounds)
	// list[len(list)]
	// list[-4] = 2
	fmt.Println(list)

	// a slice is basically an array with a variable length (not constant length)
	// slice declarations are a little different
	var a []int // right now, a has a special value nil
	n := 4

	// slices must be made
	a = make([]int, n)

	// we set values of arrays similarly to those of slices
	a[0] = 6
	a[2] = -33

	fmt.Println(a)

	// one-line declarations are used in practice
	b := make([]int, n+2)
	fmt.Println(b)

	fmt.Println(FactorialArray(6))

	var c [6]int

	d := make([]int, 6)

	ChangeFirstElementArray(c)
	ChangeFirstelementSlice(d)

	fmt.Println("c is now", c)
	fmt.Println("d is now", d)

	fmt.Println(MinIntegers(3, 6, 7, 1, 5, -43))
}

// variadic functions take a variable number of inputs

// MinIntegers takes as input an arbitrary number of integers and returns their minimum value.
func MinIntegers(numbers ...int) int {
	if len(numbers) == 0 {
		panic("Error: empty slice given to MinIntegers.")
	}
	
	return MinIntegerArray(numbers)
}

// MinIntegerArray take as input a slice of integers and returns the minimum value in that slice.
func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: empty slice given to MinIntegerArray.")
	}
	
	var m int // will store the minimum value

	// range over list, and if we are at the 0-th element OR we find something smaller than m, update m
	for i, val := range list { // this is equivalent to for i := 0; i < len(list); i++
		if i == 0 || val < m {
			m = val
		}
	}

	return m
}

func ChangeFirstElementArray(a [6]int) {
	a[0] = 1 // copy of input array gets edited
	// copy is destroyed
}

func ChangeFirstelementSlice(a []int) {
	a[0] = 1
	// when you pass in a slice to a fucntion, you get access to that underlying array
}

// FactorialArray takes as input an integer n, and it returns an array of length n+1 whose k-th element is k!
func FactorialArray(n int) []int {
	if n < 0 {
		panic("Error: negative input given to FactorialArray.")
	}

	fact := make([]int, n+1)

	// recall: 0! = 1
	fact[0] = 1

	//range k from 1 to n, and use the fact that k! = k*(k-1)!
	for k := 1; k <= n; k++ {
		fact[k] = k*fact[k-1]
	}

	return fact
}
