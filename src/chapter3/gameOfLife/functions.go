package main

// PlayGameOfLife takes as input an initial game board and a number of generations.
// It returns a slice of game boards of length numGens+1 corresponding to playing
// the Game of Life over numGens generations, starting with the initial board.
func PlayGameOfLife(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	boards[0] = initialBoard

	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(boards[i-1])
	}

	return boards
}

// UpdateBoard takes as input a GameBoard.
// It returns the board resulting from playing the Game of Life for one generation.
func UpdateBoard(currentBoard GameBoard) GameBoard {
	// first, create new board corresponding to the new generation
	numRows := CountRows(currentBoard)
	numCols := CountColumns(currentBoard)
	newBoard := InitializeBoard(numRows, numCols)

	// range through all cells of currentBoard and update each one.
	for r := range currentBoard {
		// range over values in currentBoard[r], the current row
		for c := range currentBoard[r] {
			// curr value is currentBoard[r][c]
			newBoard[r][c] = UpdateCell(currentBoard, r, c)
		}
	}

	return newBoard
}

// InitializeBoard takes as input a number of rows and a number of columns.
// It returns a numRows x numCols GameBoard object.
func InitializeBoard(numRows, numCols int) GameBoard {
	b := make(GameBoard, numRows)
	for i := range b {
		b[i] = make([]bool, numCols)
	}

	return b
}

// UpdateCell takes a GameBoard and row/col indices r and c, and it returns the state of the board at these row/col indices is in the next generation
func UpdateCell(board GameBoard, r, c int) bool {
	numNeighbors := CountLiveNeighbors(board, r, c)

	// now it's just a matter of consulting the rules.
	if board[r][c] == true { // alive
		if numNeighbors == 2 || numNeighbors == 3 {
			return true // staying alive
		} else {
			return false // dying out
		}
	} else { // dead
		if numNeighbors == 3 { // zombie!
			return true
		} else { //RIP
			return false
		}
	}
}

// CountLiveNeighbors takes as input a GameBoard board along with row and column indices r, c.
// It returns the sum of live neighbors of board[r][c], not including cells that fall off the boundaries of the board.
func CountLiveNeighbors(board GameBoard, r, c int) int {
	count := 0

	// range over all cells in the neighborhood
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			// ensure (i, j) isn't (r, c), is on board, and is true
			if (i != r || j != c) && InField(board, i, j) && board[i][j] {
				// we found a live one!
				count++
			}
		}
	}

	return count
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
