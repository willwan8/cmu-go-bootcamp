package main

import (
	"fmt"
)

func main() {
	fmt.Println("The minimum of 3 and 4 is", Min2(3, 4))

	fmt.Println(WhichIsGreater(3, 5)) // should be -1
	fmt.Println(WhichIsGreater(42, 42)) // should be 0
	fmt.Println(WhichIsGreater(-2, -7)) // should be 1
}

// Min2 takes two integers as input and returns their minimum.
func Min2(a, b int) int {
	if a < b {
		return a
	} else { // the else statement needs to immediately follow after the closing bracket of if block | in this case we don't need this else statement
		// b must be smaller (or they are equal)
		return b
	}
}

// PositiveDifference takes as input two integers.
// It returns the absolute value of the difference between these integers.
func PositiveDifference(a, b int) int {
	var c int 
	if a == b {
		return 0
	} else if  a > b {
		c = a - b
	} else {
		// can safely infer that b > a
		c = b - a
	}
	return c
}

// SameSign takes as input two integers.
// It returns true if the two integers have the same sign and false if they have different signs.
func SameSign(x, y int) bool {
	// I'm assuming that 0 can be considered both positive and negative
	// x and y have the same sign when x * y is >= 0
	if x * y >= 0 {
		return true
	} else { // I know that x * y < 0
		return false
	}

	// one line function: return x * y >= 0 (this expression will return a boolean value)
}

// WhichIsGreater compares two input integers and returns...
// 1 if the first input is larger, 
// -1 if the second input is larger, 
// and 0 if they are equal.
func WhichIsGreater(x, y int) int {
	//we need three cases (now we utilize else if)
	if x == y {
		return 0
	} else if x > y {
		return 1
	} else {
		// if we make it into this else block, we know that y > x
		return -1
	}
}

// index of comparison operators
/*
> : more than
< : less than
>= : greater than or equal to
<= : less than or equal to
== : equal to
!= : not equal to
! : "not". For example, if x is Boolean, then !x is false when x is true and true when x is false
*/