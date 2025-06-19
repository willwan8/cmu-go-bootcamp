package main

import "fmt"

func main() {
	fmt.Println("Two-dimensional arrays (and slices).")

	var a [7][4]int // declares a 7 x 4 array with all default values
	// var a [7]([4]int) // alternate declaration; conceptualizes the array as an array of length 7 where each element is an array of length 4

	a[1][2] = 19
	a[0][0] = 42
	a[6][3] = 100

	fmt.Println(a)

	fmt.Println(a[1]) // prints the second row of a

	fmt.Println("Number of rows:", len(a))       // prints 7
	fmt.Println("Number of columns:", len(a[0])) // prints 4

	// two-dimensional slices
	numRows := 4
	board := make([][]bool, numRows) // conceptualization = board = make([]([]bool), numRows) // a 2D slice of boolean variables is a 1D slice whose elements are all 1D slices of boolean variables
	// doing this ^ will result in each element of this slice being an empty slice of boolean variables (every row has length 0)

	// make all the rows of board
	for row := range board {
		board[row] = make([]bool, row)
	}

	// appending false to the end of each row
	for row := range board {
		board[row] = append(board[row], false)
	}

	// append a new row to the board
	newRow := make([]bool, 5)
	board = append(board, newRow)

	fmt.Println(board)

	SetFirstElementToTrue(board)

	fmt.Println(board) // slices are pass-by-reference, so changes made to board inside the function will still be applied outside the function as well

	rPentomino := make([][]bool, 5)
	for row := range rPentomino {
		rPentomino[row] = make([]bool, 5)
	}

	rPentomino[1][2] = true
	rPentomino[1][3] = true
	rPentomino[2][1] = true
	rPentomino[2][2] = true
	rPentomino[3][2] = true

	PrintBoard(rPentomino)
}

// SetFirstElementToTrue takes as input a two-dimensional slice of boolean variables.
// It sets the top left element equal to true.
func SetFirstElementToTrue(board [][]bool) {
	if len(board) == 0 || len(board[0]) == 0 {
		panic("no")
	}
	board[0][0] = true
}

// PrintBoard takes as input a two-dimensional slice of booleans representing a GameBoard.
// It returns nothing but it prints the GameBoard to the console, one line at a time.
func PrintBoard(board [][]bool) {
	for r := range board {
		PrintRow(board[r])
	}
}

// PrintRow takes a slice of boolean variables as input.
// It prints all of the cells in this slice one at a time, then prints a new line.
func PrintRow(row []bool) {
	for _, val := range row {
		PrintCell(val)
	}
	fmt.Println()
}

// PrintCell takes a boolean value as input
// It prints 😍 if the value is true (alive), and 💀 if it is false (dead).
func PrintCell(value bool) {
	if value {
		fmt.Print("😍")
	} else {
		fmt.Print("💀")
	}
}
