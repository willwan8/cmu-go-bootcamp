package main

import "strconv"

// PlayAutomaton takes as input a GameBoard, an integer numGens,
// a neighborhood type as a string, and a map of strings to integers representing rules.
// It returns a slice of game boards of length numGens+1 corresponding to simulating the automaton
// specified by the rules over numGens generations, starting with the initial board.
func PlayAutomaton(initialBoard GameBoard, numGens int, neighborhoodType string, rules map[string]int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1], neighborhoodType, rules)
	}

	return boards
}

// UpdateBoard takes as input a GameBoard, a neighborhood type, and a rule map.
// It returns the board resulting from simulating the automaton for one generation according by the rules and neighborhood type.
func UpdateBoard(currBoard GameBoard, neighborhoodType string, rules map[string]int) GameBoard {
	// first, create new board corresponding to the next generation. All cells will have state 0
	numRows := CountRows(currBoard)
	numCols := CountColumns(currBoard)
	newBoard := InitializeBoard(numRows, numCols)

	// range through all cells of currBoard and update each one into newBoard.
	for r := range currBoard {
		// range over values in currBoard[r]
		for c := 0; c < numCols; c++ {
			// curr value is currBoard[r][c]
			newBoard[r][c] = UpdateCell(currBoard, r, c, neighborhoodType, rules)
		}
	}

	return newBoard
}

// UpdateCell takes a GameBoard along with row and column indices, a neighborhood type, and a rule map.
// It returns the state of the cell at this row and column in the next generation of the  update cell state at given row and col indices.
func UpdateCell(board GameBoard, r, c int, neighborhoodType string, rules map[string]int) int {
	// range through rules and look for a match ... very different for the two neighborhood types
	neighborhood := NeighborhoodToString(board, r, c, neighborhoodType)
	rule, exists := rules[neighborhood]
	if exists {
		return rule
	}
	// default rule: 0
	return 0
}

// NeighborhoodToString takes as input a GameBoard, row and column indices, and a neighborhood type as a string.
// It returns a string formed of the central square followed by its neighbors according to the neighborhood type indicated.
func NeighborhoodToString(currentBoard GameBoard, r int, c int, neighborhoodType string) string {
	// Initialize the neighborhood string with the current cell value
	neighborhood := strconv.Itoa(currentBoard[r][c])

	// Initialize the offsets based on the neighborhood type
	var offsets [][2]int
	if neighborhoodType == "Moore" {
		offsets = [][2]int{
			{-1, -1}, {-1, 0}, {-1, 1},
			{0, 1}, {1, 1}, {1, 0},
			{1, -1}, {0, -1},
		}
	} else if neighborhoodType == "vonNeumann" {
		offsets = [][2]int{
			{-1, 0}, {0, 1}, {1, 0}, {0, -1},
		}
	}

	// Iterate over the offsets to construct the neighborhood string
	for _, offset := range offsets {
		x := offset[0]
		y := offset[1]
		if InField(currentBoard, r+x, c+y) {
			neighborhood += strconv.Itoa(currentBoard[r+x][c+y])
		} else { //default case
			neighborhood += "0"
		}
	}

	return neighborhood
}

// NeighborhoodToString takes as input a GameBoard, row and column indices, and a neighborhood type as a string.
// It returns a string formed of the central square followed by its neighbors according to the neighborhood type indicated.

// InitializeBoard takes as input a number of rows and a number of columns.
// It returns a numRows x numCols GameBoard object with all values equal to zero.
func InitializeBoard(numRows, numCols int) GameBoard {
	b := make(GameBoard, numRows)
	for i := range b {
		b[i] = make([]int, numCols)
	}

	return b
}

// CountRows takes as input a GameBoard and returns the number of rows
// in the board, assuming that the board is rectangular.
func CountRows(board GameBoard) int {
	return len(board)
}

// CountColumns takes as input a GameBoard and returns the number of columns
// in the board, assuming that the board is rectangular.
func CountColumns(board GameBoard) int {
	// assert that we have a rectangular board
	AssertRectangular(board)

	// we can assume board is rectangular now
	// so number of columns is # of elements in 0-th row
	return len(board[0])
}

// AssertRectangular takes as input a GameBoard.
// If returns nothing. If the board has no rows, or is not rectangular, it panics.
func AssertRectangular(board GameBoard) {
	if len(board) == 0 {
		panic("Error: no rows in GameBoard.")
	}

	firstRowLength := len(board[0])

	//range over rows and make sure that they have the same length as first row
	for row := 1; row < len(board); row++ {
		if len(board[row]) != firstRowLength {
			panic("Error: GameBoard is not rectangular.")
		}
	}
}

// InField takes a GameBoard board as well as row and col indices (i, j).
// It returns true if board[i][j] is in the board and false otherwise.
func InField(board GameBoard, i, j int) bool {
	if i < 0 || j < 0 {
		return false
	}
	if i >= CountRows(board) || j >= CountColumns(board) {
		return false
	}
	// if we survive to here, then we are on the board
	return true
}
