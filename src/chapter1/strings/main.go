package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Strings.")

	fmt.Println(string('A'))
	fmt.Println(string(45)) // doesn't actually print 45 since we need to convert 45 to a string

	fmt.Println(strconv.Itoa(45))

	j, err := strconv.Atoi("23") // the function will return j (desired value) plus an error variable

	// the conversion is OK when error variable is equal to nil
	if err != nil {
		// a problem happened
		panic(err)
	}

	fmt.Println(j)

	pi, err2 := strconv.ParseFloat("3.14", 64) // the 64 corresponds to the float type (in this case, float64)

	if err2 != nil {
		// a problem happened
		panic(err2)
	}

	fmt.Println("The value of pi is", pi)

	var s string = "Hi"
	t := "Lovers"
	// concatenate these strings with + operation
	u := s+t
	fmt.Println(u)
	// strings are conceptually arrays of symbols of type byte, so we can access (view) the elements of a string like an array or slice
	// you cannot change individual elements of a string (cannot change symbols at specfic indexes)
	fmt.Println("The first symbol of u is", string(u[0])) // we need to convert the element to a string since the element is an integer corresponding to a symbol
	fmt.Println("The final symbol of u is", string(u[len(u)-1]))

	if t[2] == 'v' { // comparisons are case sensitive
		fmt.Println("The symbol at position 2 of t is v.")
	}

	dna := "ACCGAT"
	fmt.Println(Complement(dna)) // should print "TGGCTA"
	fmt.Println(Reverse(dna)) // should print "TAGCCA"
	fmt.Println(ReverseComplement(dna)) // should print "ATCGGT"
}

// Reversecomplement takes as input a string pattern of DNA symbols.
// It returns thte reverse complement of the string.
func ReverseComplement(pattern string) string {
	return Reverse(Complement(pattern))
}

// Complement takes as input a string dna of DNA symbols.
// It returns the string formed by complementing each position of the input string ('A' <-> 'T', 'C' <-> 'G')
func Complement(dna string) string {
	dna2 := make([]byte, len(dna))
	for i, symbol := range dna {
		switch symbol { // a switch statement is essentially a lot of if statements (cases)
			case 'A':
				dna2[i] += 'T'
			case 'C':
				dna2[i] += 'G'
			case 'G':
				dna2[i] += 'C'
			case 'T':
				dna2[i] += 'A'
			default:
				panic("Invalid symbol given to Complement().")
		}
	}
	return string(dna2) // converts slice of bytes to a string
}

// Reverse take sas input a string pattern.
// It returns the string formed by reversing the positions of all symbols in pattern.
func Reverse(pattern string) string {
	rev := make([]byte, len(pattern))

	n := len(pattern)

	for i := range pattern {
		rev[i] = pattern[n-1-i]
	}

	return string(rev)
}

/*
Algorithms to implement:

ReverseComplement(pattern)
	return Reverse(Complement(pattern))
*/