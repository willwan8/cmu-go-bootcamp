package main

import (
	"os"      // for reading from files
	"strconv" // our old friend string conversion package
	"strings" // for working with strings
)

// ReadElectoralVotes processes the number of electoral votes for each state
// Input: a filename string.
// Output: a map that associates each state name (string) to an unsigned integer corresponding to its number of Electoral College Votes.
func ReadElectoralVotes(filename string) map[string]uint {
	electoralVotes := make(map[string]uint)

	// read in the file contents
	fileContents, err := os.ReadFile(filename)
	// fileContents is a slice of bytes

	Check(err)

	giantString := string(fileContents)

	// let's split the string into lines
	lines := strings.Split(giantString, "\n") // \n means new line

	// range over lines, parse each line, and add values to our map
	for _, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")
		// lineElements has two items: the state name and the number of electoral votes (as a string)
		stateName := lineElements[0]

		// parse the number of electoral votes
		numVotes, err := strconv.Atoi(lineElements[1])
		Check(err)

		// convert to uint and place this into map
		electoralVotes[stateName] = uint(numVotes)

	}

	return electoralVotes
}

// ReadPollingData parses polling percentages fro ma file.
// Input: a filename string.
// Output: a map of state names (strings) to floats corresponding to the percentages for candidate 1.
func ReadPollingData(filename string) map[string]float64 {
	candidate1Percentages := make(map[string]float64)

	// read in the file contents
	fileContents, err := os.ReadFile(filename)
	// fileContents is a slice of bytes

	Check(err)
	//convert to string
	giantString := string(fileContents)

	// let's split the string into lines
	lines := strings.Split(giantString, "\n") // \n means new line

	// range over each line of the file and parse in the data
	for _, currentLine := range lines {
		// split the current line into its elements
		lineElements := strings.Split(currentLine, ",")
		// lineElements has there items (state name and two percentages)
		stateName := lineElements[0]
		percentage1, err := strconv.ParseFloat(lineElements[1], 64)

		Check(err)

		// normalize percentage (divide by 100) and set the appropiate map value
		candidate1Percentages[stateName] = percentage1 / 100.0
	}

	return candidate1Percentages
}

// Check takes as input a variable of type error.
// If the variable is not nil, it panics
func Check(e error) {
	if e != nil {
		panic(e)
	}
}
