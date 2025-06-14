package main

func main() {
	// HW 4.1
	/* fmt.Println("Homework problem 4.1: Richness")
	no checks for this one */
	/* fmt.Println("Homework problem 4.1: SumOfValues")
	no checks for this one */

	// HW 4.2
	/* fmt.Println("Homework problem 4.2: SumOfMinima")
	no checks for this one */
	/* fmt.Println("Homework problem 4.2: BrayCurtisDistance")
	no checks for this one */
	/* fmt.Println("Homework problem 4.2: SumOfMaxima")
	no checks for this one */
	/* fmt.Println("Homework problem 4.2: JaccardDistance")
	no checks for this one */
}

// Homework 4.1
// Richness takes a frequency table, sample, which maps strings to their number of occurrences, and returns the richness of the sample (number of elements with positive counts)
func Richness(sample map[string]int) int {
	count := 0
	for _, val := range sample {
		if val > 0 {
			count++
		}
	}

	return count
}

// SumOfValues takes a frequency table, sample, which maps strings to integers, and returns the sum of values in sample
func SumOfValues(sample map[string]int) int {
	sum := 0
	for _, val := range sample {
		sum += val
	}

	return sum
}

// SimpsonsIndex takes a frequency table, sample, which maps strings to integers, and returns the Simpson's index of sample
// (1-sum[n(n-1)/N(N-1)] where n is the number of individuals of a species and N is the total number of individuals of all species
func SimpsonsIndex(sample map[string]int) float64 {
	simpsonsIndex := 0.0
	N := SumOfValues(sample) // finds total number of values out of the sample

	for _, val := range sample {
		simpsonsIndex += (float64(val) / float64(N)) * (float64(val) / float64(N)) // equal to (n/N)^2
	}

	return simpsonsIndex
}

// Homework 4.2
// SumOfMinima takes as input two frequency tables, sample1 and sample2, and returns the sum of corresponding minimum values of sample1 and sample2
func SumOfMinima(sample1, sample2 map[string]int) int {
	sum := 0
	alreadyViewed := make(map[string]bool)

	for key := range sample1 { // access keys to compare rows
		sum += Min2(sample1[key], sample2[key])
		alreadyViewed[key] = true
	}
	// to ensure that we cover all different "species" in each table
	for key := range sample2 { // access keys to compare rows
		if !alreadyViewed[key] {
			sum += Min2(sample1[key], sample2[key])
		}
	}

	return sum
}

// Subroutine for SumOfMinima out of convenience
// Min2 takes in two integers, x and y, and returns whichever one has the minimum value
func Min2(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// BrayCurtisDistance takes as input two frequency tables, sample1 and sample2, and returns the Bray-Curtis distance between sample1 and sample2
// the Bray-Curtis distance is 1 minus the sum of the minimum count of each species in each row divided by the average of the total values
func BrayCurtisDistance(sample1, sample2 map[string]int) float64 {
	brayCurtisDistance := float64(SumOfMinima(sample1, sample2)) / ((float64(SumOfValues(sample1) + SumOfValues(sample2))) / 2.0)

	return 1 - brayCurtisDistance
}

// SumOfMaxima takes as input two frequency tables, sample1 and sample2, and returns the sum of corresponding maximum values of sample1 and sample2
func SumOfMaxima(sample1, sample2 map[string]int) int {
	sum := 0
	alreadyViewed := make(map[string]bool)

	for key := range sample1 { // access keys to compare rows
		sum += Max2(sample1[key], sample2[key])
		alreadyViewed[key] = true
	}
	// to ensure that we cover all different "species" in each table
	for key := range sample2 { // access keys to compare rows
		if !alreadyViewed[key] {
			sum += Max2(sample1[key], sample2[key])
		}
	}

	return sum
}

// Subroutine for SubOfMaxima out of convenience
// Max2 takes in two integers, x and y, and returns whichever one has the maximum value
func Max2(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// JaccardDistance takes as input two frequency tables, sample1 and sample2, and returns the Jaccard distance between sample1 and sample2
// the Jaccard distance is 1 minus the sum of the minimum count of each species in each row divided by the maximum count of each species in each row
func JaccardDistance(sample1, sample2 map[string]int) float64 {
	jaccardDistance := float64(SumOfMinima(sample1, sample2)) / float64(SumOfMaxima(sample1, sample2))

	return 1 - jaccardDistance
}
