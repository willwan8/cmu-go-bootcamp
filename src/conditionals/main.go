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