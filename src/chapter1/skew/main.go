package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// MinimumSkew
// Input: a (DNA) string genome
// Output: a slice of integers corresponding to the indices in the genome where the skew array attain a minimum value
func MinimumSkew(genome string) []int {
	indices := make([]int, 0)

	// make the skew ar ray
	array := SkewArray(genome)

	m := MinIntegerArray(array)

	for i, val := range array {
		if val == m {
			// I found an index! :)
			indices = append(indices, i)
		}
	}

	return indices
}

// MinIntegerArray
// Input: an array, list, of integers
// Output: the minimum integer value in list
func MinIntegerArray(list []int) int {
	if len(list) == 0 {
		panic("Error: empty list.")
	}
	m := list[0]

	for _, val := range list {
		if val < m {
			m = val
		}
	}
	
	return m
}

// SkewArray
// Input: a (DNA) string genome
// Output: a slice of integers corresponding to the "G-C skew" of the genome at each position
func SkewArray(genome string) []int {
	n := len(genome)

	array := make([]int, n+1)
	array[0] = 0 // redundant since all elements of the array are already set to 0

	// range over remaining values and set i-th value of array
	for i := 1; i < n+1; i++ {
		array[i] = array[i-1] + Skew(genome[i-1])
	}

	return array
}

// Skew
// Input: a symbol
// Output: 1 (if symbol is G), -1 (if symbol is C), and 0 otherwise.
func Skew(symbol byte) int {
	if symbol == 'G' {
		return 1
	} else if symbol == 'C' {
		return -1
	}
	// default
	return 0
}

func main() {
	fmt.Println("The skew array.")

	url := "https://bioinformaticsalgorithms.com/data/realdatasets/Replication/E_coli.txt"

	// grab response from site
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close() // close connection later

	// status OK?
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Received non-OK status: %v", resp.Status)
	}

	// access slice of symbols in file
	genomeSymbols, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// convert genome to string
	genome := string(genomeSymbols)
	fmt.Println("Genome read. It has", len(genome), "total nucleotides.")

	EcoliSkewArray := SkewArray(genome)
	minSkewPositions := MinimumSkew(genome)

	firstPosition := minSkewPositions[0]

	fmt.Println("The minimum skew of", EcoliSkewArray[firstPosition], "occurs at positions", minSkewPositions)
}