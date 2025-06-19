package main

import (
	"log"     // needed for error reporting
	"os"      // needed for reading in from file
	"strings" // needed for manipulating data parsed from file
)

// ReadBoardFromFile takes a filename as a string and reads in the data provided
// in this file, returning a game board.
func ReadBoardFromFile(filename string) GameBoard {
	// try reading in the file
	fileContents, err := os.ReadFile(filename)

	Check(err)

	// parse out the file contents as a slice of strings, one line per element
	giantString := string(fileContents)
	trimmedGiantString := strings.TrimSpace((giantString))
	lines := strings.Split(trimmedGiantString, "\n") // the number of lines = the number of rows in our GameBoard

	// we know how many rows the GameBoard will have
	board := make(GameBoard, len(lines))

	// next, parse the data on each line and add to the GameBoard
	for i, currentLine := range lines {
		lineElements := strings.Split(currentLine, ",")

		// use a subroutine to set the values of the current row
		board[i] = SetRowValues(lineElements)
	}

	return board
}

// Check takes as input a variable of error type.
// If the variable has any value other than nil, it panics.
func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// SetRowValues takes as input a slice of strings, where each element is either a "0" or "1".
// It returns a slice of boolean values of the same length, encoding "0" as false and "1" as true.
func SetRowValues(lineElements []string) []bool {
	currentRow := make([]bool, len(lineElements))

	// range through the current line elements and set values
	for j, val := range lineElements {
		if val == "0" {
			currentRow[j] = false // technically this is not necessary (default values are already false)
		} else if val == "1" {
			currentRow[j] = true
		} else {
			log.Fatal("Error: invalid entry in board file.")
		}
	}

	return currentRow
}
