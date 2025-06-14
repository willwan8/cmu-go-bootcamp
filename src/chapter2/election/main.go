package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Let's simulate an election!")

	electoralVoteFile := "data/electoralVotes.csv"
	pollFile := "data/debates.csv"

	// now, read them in and store as maps
	electoralVotes := ReadElectoralVotes(electoralVoteFile)
	polls := ReadPollingData(pollFile)

	numTrials := 1000000
	marginOfError := 0.1

	probability1, probability2, probabilityTie := SimulateMultipleElections(polls, electoralVotes, numTrials, marginOfError)

	fmt.Println("Probability of candidate 1 winning:", probability1)

	fmt.Println("Probability of candidate 2 winning:", probability2)

	fmt.Println("Probability of tie:", probabilityTie)

}

// SimulateMultipleelections runs a Monte Carlo simulation with multiple trials to simulate a presidential election.
// Input: polling data as a map of states to percentages for candidate 1, a map of state names to Electoral College votes,
// an integer corresponding to the number of trails to run, and a decimal margin of error.
// Output: Three probabilities corresponding to the likelihood of each of two candidates winning or a tie.
func SimulateMultipleElections(polls map[string]float64, electoralVotes map[string]uint, numTrials int, marginOfError float64) (float64, float64, float64) {
	winCount1 := 0
	winCount2 := 0
	tieCount := 0

	// simulate a single election n times and update count each time
	for i := 0; i < numTrials; i++ {
		// simulate one election
		votes1, votes2 := SimulateOneElection(polls, electoralVotes, marginOfError)

		// who won?
		if votes1 > votes2 {
			winCount1++
		} else if votes2 > votes1 {
			winCount2++
		} else {
			// dreaded die!
			tieCount++
		}
	}

	// divide number of wins by number of trails
	probability1 := float64(winCount1) / float64(numTrials)
	probability2 := float64(winCount2) / float64(numTrials)
	probabilityTie := float64(tieCount) / float64(numTrials)

	return probability1, probability2, probabilityTie
}

// SimulateOneElection simulates a single election from polling data.
// Input: polling data as a map of states to percentages for candidate 1, a map of state names to Electoral College votes, and a decimal margin of error.
// Output: the number of Electoral College votes in a simulated election for each of candidate 1 and candidate 2.
func SimulateOneElection(polls map[string]float64, electoralVotes map[string]uint, marginOfError float64) (uint, uint) {
	var collegeVotes1 uint
	var collegeVotes2 uint

	// range over all the states, and simulate the election in each one.
	for state, pollingValue := range polls {
		// first, let's grab the number of EC votes
		numVotes := electoralVotes[state]

		// let's adjust the polling value with some noise
		adjustedPoll := AddNoise(pollingValue, marginOfError)

		// who won the state? (based on adjusted number)
		if adjustedPoll >= 0.5 {
			collegeVotes1 += numVotes
		} else {
			collegeVotes2 += numVotes
		}
	}

	return collegeVotes1, collegeVotes2
}

// AddNoise adjusts a polling percentage based on some randomization sampled from a normal distribution.
// Input: Two decimals, a polling value, and a margin of error
// Output: A decimal corresponding to the adjusted polling value.
func AddNoise(pollingValue, marginOfError float64) float64 {
	x := rand.NormFloat64()
	// x has a 95% chance of being between -2 and 2

	x /= 2.0
	// x has a 95% chance of being between -1 and 1

	x *= marginOfError

	// x has a 95% chance of being between -marginOfError and +marginOfError
	return x + pollingValue
}
