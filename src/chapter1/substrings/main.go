package main

import (
	"fmt"
	// "strings"
)

func main() {
	fmt.Println("Substrings (and subslices)")

	s := "Hi Lovers!"

	fmt.Println(s[1:5]) // prints a substring of s starting at index 1 and ending at index 5 (exclusive) => "i Lo"
	fmt.Println(s[:7]) // prefix of a string => this is equal to s[0:7]
	fmt.Println(s[4:]) // suffix of a string => this is equal to s[4:len(s)]

	a := make([]int, 6)
	for i := range a {
		a[i] = 2*i + 1
	}
	fmt.Println(a)
	fmt.Println(a[3:5]) // prints the value from index 3 up to index 5
	fmt.Println(a[:3]) // we can also use the prefix shorthand notation
	fmt.Println(a[4:]) // we can also use the suffix shorthand notation

	pattern := "ATA"
	text := "ATATA"

	fmt.Println(PatternCount(pattern, text))
	fmt.Println(StartingIndices(pattern, text))

	//fmt.Println(strings.Count(text, pattern)) // built in function to count patterns within strings, but excludes overlaps
}

// PatternCount takes as input two strings pattern and text.
// It returns the number of times that pattern occurs as a substring of text, including overlaps.
func PatternCount(pattern, text string) int {
	return len(StartingIndices(pattern, text)) // if we find all starting indices of a pattern in StartingIndices, we can just take the length of the slice returned to find the number of patterns
	/*
	count := 0
	n := len(text)
	k := len(pattern)

	// range over text, incrementing count every time we find a pattern match
	for i := 0; i < n-k+1; i++ {
		if pattern == text[i:i+k] {
			count++
		}
	}

	return count
	*/
}

// StartingIndices takes as input two strings pattern and text.
// It returns the colletion of all starting positions of pattern as a substring of text, including overlaps.
func StartingIndices(pattern, text string) []int {
	positions := make([]int, 0)

	n := len(text)
	k := len(pattern)

	// range over text, appending current position to positions if we find a pattern match
	for i := 0; i < n-k+1; i++ {
		if pattern == text[i:i+k] {
			positions = append(positions, i)
		}
	}

	return positions
}
/*
PatternCount(pattern, text)
  count ← 0
  n ← length(text)
  k ← length(pattern)
  for every integer i between 0 and n − k
    if text[i, i + k] = pattern
      count ← count + 1
  return count

StartingIndices(pattern, text)
  positions ← array of integers of length 0
  n ← length(text)
  k ← length(pattern)
  for every integer i between 0 and n − k
    if text[i, i + k] = pattern
      positions ← append(positions, i)
  return positions
*/