package main
import (
	"fmt"
)

func main() {
	fmt.Println("Finding frequent words.")

	text := "ATCAATGATCAACGTAAGCTTCTAAGCATGATCAAGGTGCTCACACAGTTTATCCACAACCTGAGTGGATGACATCAAGATAGGTCGTTGTATCTCCTTCCTCTCGTACTCTCATGACCACGGAAAGATGATCAAGAGAGGATGATTTCTTGGCCATATCGCAATGAATACTTGTGACTTGTGCTTCCAATTGACATCTTCAGCGCCATATTGCGCTGGCCAAGGTGACGGAGCGGGATTACGAAAGCATGATCATGGCTGTTGTTCTGTTTATCTTGTTTTGACTGAGACTTGTTAGGATAGACGGTTTTTCATCACTGACTAGCCAAAGCCTTACTCTGCCTGACATCGACCGTAAATTGATAATGAATTTACATGCTTCCGCGACGATTTACCTCTTGATCATCGATCCGATTGAAGATCTTCAATTGTTAATTCTCTTGCCTCGACTCATAGCCATGATGAGCTCTTGATCATGTTTCCTTAACCCTCTATTTTTTACGGAAGAATGATCAAGCTGCTGCTCTTGATCATCGTTTC"

	k := 9

	freqMap := FrequencyTable(text, k)

	freqPatterns := FindFrequentWords(text, k)

	fmt.Println("Frequent pattern(s) found when k is", k, "are", freqPatterns)

	fmt.Println("The maximum number of occurrences is", freqMap[freqPatterns[0]])
}

// FindFrequentWords takes as input a string text and an integer k.
// It returns a slice of strings corresponding to the substring(s) of length k that occur most frequently in text.
func FindFrequentWords(text string, k int) []string {
	freqPatterns := make([]string, 0)

	freqMap := FrequencyTable(text, k)
	// once i Have the frequency talbe, find the maximum value
	max := MaxMapValue(freqMap)

	// what keys of this table reach the max value?
	for pattern, val := range freqMap {
		if val == max {
			// append!
			freqPatterns = append(freqPatterns, pattern)
		}
	}

	return freqPatterns
}

// MaxMapValue takes as input a map of strings to integers.
// It returns the maximum value found in the map
func MaxMapValue(dict map[string]int) int {
	m := 0
	firstTimeThrough := true

	// range over map and update m every time I find a value that is larger
	for _, val := range dict {
		if firstTimeThrough || val > m {
			m = val
			firstTimeThrough = false
		}
	}

	return m
}

// FrequencyTale takes as input a string text and an integer k.
// It returns the frequency table mapping each substring of text of length k to its number of occurrences.
func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)

	// range over all substrings of length k
	for i := 0; i < n-k+1; i++ {
		// grab current pattern
		pattern := text[i:i+k]

		// updating the value of freqMap associated with pattern

		/*
		_, exists := freqMap[pattern]
		if !exists { // equal to exists == false
			// pattern has not been encountered
			freqMap[pattern] = 1
		} else {
			// pattern has been encountered
			freqMap[pattern]++
		}
		*/

		freqMap[pattern]++
		// if pattern is a key, this is what we want
		// if it's not, freqMap[pattern] gets created, gets set equal to zero, then incremented

	}

	return freqMap
}