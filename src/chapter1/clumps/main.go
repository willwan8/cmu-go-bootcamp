package main

import (
	"fmt"
	"net/http" // for accessing URLs
	"io" // needed to read from (and write to) files
	"log" // needed for log files (errors)
	"time"
)

func main() {
	fmt.Println("Clumps.")

	/*
	text := "AAAACGTCGAAAAA"
	
	k := 2
	L := 4
	t := 2

	fmt.Println(FindClumps(text, k, L, t))
	*/

	// Getting and reading genome files from urls
	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// reponse was OK
	
	// close the connection after you're done
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK status: %v", resp.Status)
	}

	genomeSymbols, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	// we now have a slice of symbols associated with the genome

	fmt.Println("The number of nucleotides in E.coli genome is", len(genomeSymbols))

	EcoliGenome := string(genomeSymbols)

	k := 9
	L := 500
	t := 3

	//clumps := FindClumps(EcoliGenome, k, L, t)
	/*
	start := time.Now()
	FindClumps(EcoliGenome, k, L, t)
	elapsed := timeSince(start)
	log.Printf("FindClumps took %s", elapsed)
	*/

	start2 := time.Now()
	clumps := FindClumpsFaster(EcoliGenome, k, L, t)
	elapsed2 := time.Since(start2)
	log.Printf("FindClumpsFaster took %s", elapsed2)
	
	fmt.Println("Found", len(clumps), "total patterns occuring as clumps.")
}

// FindClumpsFaster takes as input a string text, and integers k, L, and t.
// It returns a slice of strings representing all substrings of length k that appear at least t times in a "window" of length L in the string text.
func FindClumpsFaster(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	n := len(text)

	// map to track whether I have added a string to patterns yet
	foundPatterns := make(map[string]bool)

	firstWindow := text[:L]

	freqMap := FrequencyTable(firstWindow, k)

	// append any patterns we find to patterns slice
	for s, freq := range freqMap {
		if freq >= t {
			patterns = append(patterns, s)
			foundPatterns[s] = true
		}
	}

	// range over all remaining possible windows of text
	for i := 1; i < n-L+1; i++ {
		// decrease by 1 the value associated with the first substring of length k in the preceding window
		oldPattern := text[i-1:i-1+k]
		freqMap[oldPattern]--

		// clean up the map if the frequency of oldPattern was 1
		if freqMap[oldPattern] == 0 {
			delete(freqMap, oldPattern)
		}

		// add the pattern from the end of the current window
		newPattern := text[i+L-k:i+L]
		freqMap[newPattern]++

		// I have updated the frequency map :)
		for s, freq := range freqMap {
			if freq >= t && !foundPatterns[s] {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}

	}

	return patterns
}

// FindClumps takes as input a string text, and integers k, L, and t.
// It returns a slice of strings representing all substrings of length k that appear at least t times in a "window" of length L in the string text.
func FindClumps(text string, k, L, t int) []string {
	patterns := make([]string, 0)
	n := len(text)

	// map to track whether I have added a string to patterns yet
	foundPatterns := make(map[string]bool)

	// range over all possible windows of text
	for i := 0; i < n-L+1; i++ {
		// set the current window
		window := text[i: i+L]

		// let's make the frequency table for this window
		freqMap := FrequencyTable(window, k)

		// what occurs frequently (i.e., t or more times)?
		for s, val := range freqMap {
			// append s to patterns if s occurs frequently and s doesn't already appear in patterns
			if val >= t && foundPatterns[s] == false {
				patterns = append(patterns, s)
				foundPatterns[s] = true
			}
		}
	}

	return patterns
}

// Contains takes as input a slice of strings and a single string s.
// It returns true if s is contained in the slice, and false otherwise.
/*
func Contains(paterns []string, s string) bool {
	for _, pattern : range patterns {
		if pattern == s {
			return true
		}
	}
	// default: return false
	return false
}
*/

// FrequencyTale takes as input a string text and an integer k.
// It returns the frequency table mapping each substring of text of length k to its number of occurrences.
func FrequencyTable(text string, k int) map[string]int {
	freqMap := make(map[string]int)
	n := len(text)

	// range over all substrings of length k
	for i := 0; i < n-k+1; i++ {
		// grab current pattern
		pattern := text[i:i+k]

		freqMap[pattern]++
		// if pattern is a key, this is what we want
		// if it's not, freqMap[pattern] gets created, gets set equal to zero, then incremented
	}

	return freqMap
}