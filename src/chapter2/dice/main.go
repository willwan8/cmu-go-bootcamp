package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Rolling dice.")

	//rand.Seed(1) // we set the position where the pseudorandom algorithm starts
	// rand.Seed(time.Now().UnixNano()) // this is now built in (seeding random values to get random numbers from algorithms)

	fmt.Println(rand.Int())     // print integer
	fmt.Println(rand.Intn(10))  // print integer between 0 and 9
	fmt.Println(rand.Float64()) // print decimal in range [0, 1)

	fmt.Println(SumDice(2))
}

// RollDie
// Input: none
// Output: a pseudorandom integer between 1 and 6, inclusively.
func RollDie() int {
	return rand.Intn(6) + 1
}

// SumTwoDice
// Input: none
// Output: the simulated sum of two dice (between 2 and 12).
func SumTwoDice() int {
	return RollDie() + RollDie()
	/*
		roll := rand.Float64()
		if roll < 1.0/36.0 {
			return 2
		} else if roll < 3.0/36.0 { // we know that roll > 1/36
			return 3
		} else if roll < 6.0/36.0 {
			return 4
		} // etc.
	*/
}

// SumDice simulates the process of summing n dice
// Input: an integer numDice
// Output: the sum of numDice simulated dice
func SumDice(numDice int) int {
	sum := 0

	for i := 0; i < numDice; i++ {
		sum += RollDie()
	}

	return sum
}
