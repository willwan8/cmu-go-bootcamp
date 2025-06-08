package main

import (
	"fmt"
	"time"
	"log"
)

func main() {
	fmt.Println("Implementing two GCD algorithms.")

	fmt.Println(TrivialGCD(63, 42))

	fmt.Println(EuclidGCD(63, 42))

	x := 37820268014
	y := 27314794329

	// timing TrivialGCD
	start := time.Now() // start the stopwatch
	TrivialGCD(x, y)
	elapsed := time.Since(start) // stops the stopwatch
	log.Printf("TrivialGCD took %s", elapsed) // print to console in a pretty way

	// timing EuclidGCD
	start2 := time.Now()
	EuclidGCD(x, y)
	elapsed2 := time.Since(start2)
	log.Printf("EuclidGCD took %s", elapsed2)
}

func Min2(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// TrivialGCD takes as input two integers a and b and returns their GCD, 
// by applying a trivial algorithm attempting each possible divisor of a and b up to their minimum.
func TrivialGCD(a, b int) int {
	d := 1

	m := Min2(a, b)
	for p := 1; p <= m; p++ {
		// p is a divisor of a if a % p == 0
		// p is a divisor of b if a % p == 0
		if (a % p == 0) && (b % p == 0) {
			// if we are here, we know that p is a divisor of both
			d = p
		}
	}

	return d
}

// EuclidGCD takes as input two integers a and b and returns their GCD, following Euclid's algorithm.
func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			// we know that b > a
			b -= a
		}
	}
	return a // or b, because they are equal
}

/*
//Pseudocode functions in this code along
TrivialGCD(a, b)
  d ← 1
  m ← Min2(a,b)
  for every integer p between 1 and m
    if p is a divisor of both a and b
        d ← p
  return d

EuclidGCD(a, b)
  while a ≠ b
      if a > b
        a ← a − b
      else
        b ← b − a
  return a 
*/

/*
// Truth tables
// AND is represented by &&; OR is represented by ||
x				y				x && y	x || y
true		true		true		true
true		false		false		true
false		true		false		true
false		false		false		false
*/