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

	numTrials := 10000000

	fmt.Println("Estimated house edge with", numTrials, "trials is:", ComputeCrapsHouseEdge(numTrials))
}

// PlayCrapsOnce simulates one game of craps.
// Input: none
// Output: true if the game is a win and false if it's a loss
func PlayCrapsOnce() bool {
	firstRoll := SumDice(2)
	if firstRoll == 7 || firstRoll == 11 {
		// winner!
		return true
	} else if firstRoll == 2 || firstRoll == 3 || firstRoll == 12 {
		return false // loser! :(
	} else {
		// keep rolling until you hit a 7 or your original roll
		for { // while forever
			newRoll := SumDice(2)
			if newRoll == firstRoll {
				return true // winner!
			} else if newRoll == 7 {
				return false // loser! :(
			}
		}
	}
}

// ComputeCrapsHouseEdge estimates the "house edge" of craps over multilpe simulations.
// Input: an integer corresponding to the number of simulations.
// Output: house edge of craps (average amount won or lost over the number of simulations per unit of currency)
func ComputeCrapsHouseEdge(numTrials int) float64 {
	count := 0 // will keep track of amount wont (+) or lost (-)

	// run n trials and update count accordingly
	for i := 0; i < numTrials; i++ {
		// play the game
		outcome := PlayCrapsOnce()
		if outcome {
			// winner!
			count++
		} else {
			// loser! :(
			count--
		}
	}

	return float64(count) / float64(numTrials)
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
