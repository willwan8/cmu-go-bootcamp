package main

import (
	"chapter3/canvas" // for drawing shapes
	"image"           // for exporting our drawing to an image
)

// DrawGameBoard takes as input a (rectangular) GameBoard and an integer cellWidth.
// It returns an Image object corresponding to the GameBoard,
// where each cell in the image has width and height equal to cellWidth.
func DrawGameBoard(board GameBoard, cellWidth int) image.Image {
	// create a canvas object of the appropriate width and height (in pixels)
	height := CountRows(board) * cellWidth
	width := CountColumns(board) * cellWidth
	c := canvas.CreateNewCanvas(width, height)

	// make two colors: dark gray and white
	darkGray := canvas.MakeColor(60, 60, 60)
	white := canvas.MakeColor(255, 255, 255)

	// set the entire board as dark gray
	c.SetFillColor(darkGray)
	c.Clear()

	// fill in the colored squares
	for i := range board {
		for j := range board[i] {
			// only draw the cell if it's alive, since it would equal background color if it's dead
			if board[i][j] == true {
				c.SetFillColor(white)

				// set coordinates of the cell's top left corner
				x := float64(j*cellWidth) + float64(cellWidth)/2
				y := float64(i*cellWidth) + float64(cellWidth)/2

				scalingFactor := 0.8 // t o make circle smaller

				c.Circle(x, y, scalingFactor*float64(cellWidth)/2)
				c.Fill()
			}
		}
	}

	return c.GetImage()
}

// CountRows takes as input a GameBoard and returns the number of rows in the board
func CountRows(board GameBoard) int {
	return len(board)
}

// CountColumns takes as input a GameBoard and returns the number of columns in the board, assuming that the board is rectangular.
func CountColumns(board GameBoard) int {
	// assert that we have a rectangular board
	AssertRectangular(board)

	// we can assume board is rectangular now
	// number of columns is the number of elements in 0-th row
	return len(board[0])
}

// AssertRectangular takes as input a GameBoard.
// It returns nothing. If the board has no rows, or is not rectangular, it panics.
func AssertRectangular(board GameBoard) {
	if len(board) == 0 {
		panic("Error: no rows in GameBoard.")
	}

	firstRowLength := len(board[0])

	// range over rows and make sure that they have the same length as first row
	for row := 1; row < len(board); row++ {
		if len(board[row]) != firstRowLength {
			panic("error: GameBoard is not rectangular")
		}
	}
}

// DrawGameBoards takes as input a slice of (rectangular) Gameboards and an integer cellWidth.
// It returns a slice of Image objects corresponding to each of these GameBoards,
// where each cell in the image has width and height equal to cellWidth
func DrawGameBoards(boards []GameBoard, cellWidth int) []image.Image {
	numGenerations := len(boards)
	imageList := make([]image.Image, numGenerations)
	for i := range boards {
		imageList[i] = DrawGameBoard(boards[i], cellWidth)
	}
	return imageList
}
