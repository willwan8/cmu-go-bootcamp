package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
)

func main() {
	// HW 6.1
	fmt.Println("Homework problem 6.1: WeightedDie")
	fmt.Println(WeightedDie())
	fmt.Println("Homework problem 6.1: EstimatePi")
	fmt.Println(EstimatePi(1000000)) // Law of Large Numbers => larger # of trials = closer estimate to pi
	// HW 6.2
	fmt.Println("Homework problem 6.2: RelativelyPrime")
	fmt.Println(RelativelyPrime(1, 1), RelativelyPrime(1, 100), RelativelyPrime(2, 2), RelativelyPrime(2, 4), RelativelyPrime(6, 9), RelativelyPrime(4, 9), RelativelyPrime(100, 1024))
	fmt.Println("Homework problem 6.2: RelativelyPrimeProbability")
	fmt.Println(RelativelyPrimeProbability(1, 1, 10), RelativelyPrimeProbability(2, 2, 1), RelativelyPrimeProbability(1, 2, 100000), RelativelyPrimeProbability(3, 10, 100000), RelativelyPrimeProbability(1, 100000, 10000))
	// fmt.Println(RelativelyPrimeProbability(1, 1000000000000, 1000000000))
	// HW 6.3
	/* fmt.Println("Homework problem 6.3: HasRepeat")
	no checks for this one */
	fmt.Println("Homework problem 6.3: SimulateOneBirthdayTrial")
	fmt.Println(SimulateOneBirthdayTrial(10), SimulateOneBirthdayTrial(50), SimulateOneBirthdayTrial(100))
	fmt.Println("Homework problem 6.3: SharedBirthdayProbability")
	fmt.Println(SharedBirthdayProbability(1, 1), SharedBirthdayProbability(366, 1), SharedBirthdayProbability(11, 100000), SharedBirthdayProbability(19, 100000), SharedBirthdayProbability(40, 100000), SharedBirthdayProbability(58, 100000))
	// fmt.Println(SharedBirthdayProbability(23, 1000000))
	// HW 6.4
	fmt.Println("Homework problem 6.4: CountNumDigits")
	fmt.Println(CountNumDigits(0), CountNumDigits(-1), CountNumDigits(-498212910), CountNumDigits(155), CountNumDigits(1000000), CountNumDigits(999999))
	fmt.Println("Homework problem 6.4: SquareMiddle")
	fmt.Println(SquareMiddle(0, 4), SquareMiddle(1, 6), SquareMiddle(534243, 6), SquareMiddle(10000, 4), SquareMiddle(-1, 2), SquareMiddle(1, -5), SquareMiddle(1039, 5), SquareMiddle(5948201, 8), SquareMiddle(44, 4))
	fmt.Println("Homework problem 6.4: GenerateMiddleSquareSequence")
	fmt.Println(GenerateMiddleSquareSequence(0, 4), GenerateMiddleSquareSequence(3792, 4), GenerateMiddleSquareSequence(23, 2), GenerateMiddleSquareSequence(1600, 4), GenerateMiddleSquareSequence(8100, 4), GenerateMiddleSquareSequence(356, 6))
	/* fmt.Println("Homework problem 6.4: ComputePeriodLength")
	no checks for this one */
	// exercise for 6.4
	period10OrSmaller := 0
	for i := 1; i < 10000; i++ {
		if ComputePeriodLength(GenerateMiddleSquareSequence(i, 4)) <= 10 {
			period10OrSmaller++
		}
	}
	fmt.Println(period10OrSmaller)

	// HW 6.5
	fmt.Println("Homework problem 6.5: GenerateLinearCongruenceSequence")
	fmt.Println(GenerateLinearCongruenceSequence(1, 2, 0, 9))
	// fmt.Println(ComputePeriodLength(GenerateLinearCongruenceSequence(1, 5, 1, 8192)))

}

// Homework 6.1
// WeightedDie returns a pseudorandom integer between 1 and 6, inclusively, such that the probability of rolling a 3 is .5, and the probability of rolling any other number is .1
func WeightedDie() int {
	x := rand.Float64() // gives a value with range of [0, 1)

	if x < .5 {
		return 3
	} else if x < .6 {
		return 1
	} else if x < .7 {
		return 2
	} else if x < .8 {
		return 4
	} else if x < .9 {
		return 5
	} else {
		return 6
	}

	/* switch statement way
	switch {
	case x < .5:
		return 3
	case x < .6: // we must have passed the previous case, so this only checks values from .5 to .6 (overall, .1 probability)
		return 1
	case x < .7: // same reasoning as previous case and for other future cases
		return 2
	case x < .8:
		return 4
	case x < .9:
		return 5
	default:
		return 6 // all other cases must have failed, so this leaves a .1 probability to reach here
	}
	*/
}

// EstimatePi takes in an integer, numPoints, and returns
// an estimate of the value of pi produced by sampling numPoints points uniformly from the square of side 2 centered at the origin,
// and multiplying 4 times the fraction of these points falling in the circle of radius 1 centered at the origin
func EstimatePi(numPoints int) float64 {
	piPoint := 0 // this will keep track of the points that fall in the circle of radius 1 centered at the origin

	for i := 0; i < numPoints; i++ {
		// first, we randomly select a point in the square of side 2 centered at the origin
		// find a random x coordinate for our point that is in our square
		xCoord := rand.Float64()

		// determine whether our coordinate is positive or negative by random chance (needed so we include ALL potential coordinates)
		if !PosOrNeg() {
			xCoord = -xCoord
		}

		// find a random y coordinate for our point that is in our square
		yCoord := rand.Float64()

		// determine whether our coordinate is positive or negative by random chance (needed so we include ALL potential coordinates)
		if !PosOrNeg() {
			yCoord = -yCoord
		}

		// now we need to check if this point is in the circle of radius 1 centered at the origin
		// any point in the circle has to have a distance from the origin less than one since the circle has a radius of 1
		if DistanceFromOrigin(xCoord, yCoord) < 1 {
			piPoint++
		}

	}

	return 4.0 * (float64(piPoint) / float64(numPoints))
}

// PosOrNeg returns true if a number should be positive and false if it should be negative
func PosOrNeg() bool {
	temp := rand.Int()

	if temp%2 == 0 {
		return true
	} else {
		return false
	}
}

// DistanceFromOrigin takes in two coordinates, xCoord and yCoord, and returns their distance from the origin using the distance formula
func DistanceFromOrigin(xCoord, yCoord float64) float64 {
	return math.Sqrt((xCoord * xCoord) + (yCoord * yCoord))
}

// Homework 6.2
// RelativelyPrime takes in two integers, a and b, and returns true if a and b are relatively prime (their only common factor is 1) and false otherwise
func RelativelyPrime(a, b int) bool {
	if EuclidGCD(a, b) == 1 {
		// a and b must have a common factor of only 1 (greatest common divisor is 1)
		return true
	} else {
		// a and b must have a common factor greater than 1
		return false
	}
}

// subroutine
// EuclidGCD finds the greatest common divisor of a and b
func EuclidGCD(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

// RelativelyPrimeProbability takes three integers (lowerBound, upperBound, and numPairs) and returns an estimate of the probability that
// two randomly chosen numbers between lowerBound and upperBound are relatively prime by selecting numPairs pairs of numbers between lowerBound and upperBound inclusively
func RelativelyPrimeProbability(lowerBound, upperBound, numPairs int) float64 {
	relativelyPrime := 0
	for i := 0; i < numPairs; i++ {
		// find a random number between lowerBound and upperBound inclusively
		num1 := rand.Intn((upperBound + 1) - lowerBound) // we can add this number to lowerBound to get our desired number in the range (+1 to make this inclusive)
		num1 += lowerBound

		// find our second random number between lowerBound and upperBound inclusively
		num2 := rand.Intn((upperBound + 1) - lowerBound)
		num2 += lowerBound

		// check if this pair is relatively prime
		if RelativelyPrime(num1, num2) {
			relativelyPrime++
		}
	}

	return float64(relativelyPrime) / float64(numPairs)
}

// Homework 6.3
// HasRepeat takse an array of integers, a, and returns true if a contains a repeated value and false otherwise
func HasRepeat(a []int) bool {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				return true
			}
		}
	}
	// there must be no repeated values
	return false
}

// SimulateOneBirthdayTrial takes an integer, numPeople, and returns true or false, corresponding to whether in a room
// of numPeople randomly generated people, at least one pair of people have the same birthday
func SimulateOneBirthdayTrial(numPeople int) bool {
	// we assume that a year has 365 days, so we can randomly generate a number from 1 to 365, inclusive, to represent a date of birth
	// we can store these values in an array, and if the array has a repeated value, then we know a pair was born on the same day
	var birthdays []int

	// creating our array of randomly generated birthdays
	for i := 0; i < numPeople; i++ {
		birthday := rand.Intn(365) + 1 // we add 1 so that our range is from 1 to 365 inclusive
		birthdays = append(birthdays, birthday)
	}

	// if our array has a repeat, then we return true, else we return false
	// HasRepeat already does this, so we can just return a call to HasRepeat since HasRepeat returns a boolean
	return HasRepeat(birthdays)
}

// SharedBirthdayProbability takes in two integers, numPeople and numTrials, and returns an estimate of the probability that
// in a room of numPeople randomly gnerated people, at least one pair of people have the same birthday after running numTrials simulations
func SharedBirthdayProbability(numPeople, numTrials int) float64 {
	numOfSameBirthdayTrials := 0

	for i := 0; i < numTrials; i++ {
		if SimulateOneBirthdayTrial(numPeople) {
			numOfSameBirthdayTrials++
		}
	}

	return float64(numOfSameBirthdayTrials) / float64(numTrials)
}

// Homework 6.4
// CountNumDigits takes an integer, x, and returns the number of digits in x
func CountNumDigits(x int) int {
	if x == 0 {
		return 1
	}

	numDigits := 0

	if x < 0 { // to make x positive
		x = -x
	}

	for x > 0 {
		// we can utilize integer division to remove single digits from x and count the number of times we do this to get our digit count
		x /= 10
		numDigits++
	}

	return numDigits
}

// SquareMiddle takes integers x and numDigits and returns the result of squaring x and taking the "middle" numDigits digits in the resulting number that has 2 · numDigits total digits.
// You should add zeroes to the start of the number if needed so that it has 2 · numDigits total digits. Furthermore, your function should return -1 if numDigits is odd, if x is negative,
// if numDigits is not positive, or if the number of digits in x is greater than numDigits.
func SquareMiddle(x, numDigits int) int {
	// first, we need to check for all of the mentioned check cases
	if numDigits%2 != 0 || x < 0 || numDigits < 0 || CountNumDigits(x) > numDigits {
		return -1
	}

	// if we are here, we can proceed since we have passed all of the mentioned cases
	x *= x // squaring x

	// we want to create an array of each individual digit in our squared x
	// we can actually turn our integer into a string, get the substring we want, then turn this back into an integer to get our result
	xString := strconv.Itoa(x)

	// we add 0s in front of our digits if our length isn't at the requirement
	for len(xString) < numDigits*2 {
		xString = "0" + xString
	}

	// we can get the substring we want by starting at index numDigits/2 and ending at the length of the string minus numDigits/2
	// the length of our string should be numDigits * 2
	middleDigits := xString[numDigits/2 : len(xString)-(numDigits/2)]

	result, err := strconv.Atoi(middleDigits)

	if err != nil {
		panic(err)
	}

	return result
}

// GenerateMiddleSquareSequence takes integers seed and numDigits and returns the sequence of integers for the middle-square sequence starting with seed and using the middle numDigits digits to square each integer
func GenerateMiddleSquareSequence(seed, numDigits int) []int {
	var middleSquareSequences []int

	middleSquareSequences = append(middleSquareSequences, seed)

	index := 0

	// we continue going over the sequence until we have a repeat result
	for !HasRepeat(middleSquareSequences) {
		middleSquareSequences = append(middleSquareSequences, SquareMiddle(middleSquareSequences[index], numDigits))
		index++
	}

	return middleSquareSequences
}

// ComputePeriodLength takes an array of integers, a, and returns the period of the sequence of integers a, which is equal to one more than the number of integers between the first two repeated integers
func ComputePeriodLength(a []int) int {
	period := 0

	// we can find the period by simply subtracting the indexes at which the first two repeated integers are
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] == a[j] {
				period = j - i
			}
		}
	}

	return period
}

// Homework 6.5
// GenerateLinearCongruenceSequence takes integers seed, a, c, and m, and returns the sequence of integers produced by the linear congruential generator with the parameters a, c, and m and using seed as a starting integer.
// Note that the sequence will halt as soon as it hits a repeat, so that the first repeated element will be the final integer in the sequence
func GenerateLinearCongruenceSequence(seed, a, c, m int) []int {
	var sequence []int

	sequence = append(sequence, seed)

	index := 0
	for !HasRepeat(sequence) {
		sequence = append(sequence, (a*sequence[index]+c)%m)
		index++
	}

	return sequence
}
