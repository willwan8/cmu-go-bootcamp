package main

import (
	"fmt"
	// "image/png" // for drawing an image to PNG
	"chapter3/gameOfLife/gifhelper"
	"os" // for opening a file that will hold the image
	"strconv"
)

func main() {
	fmt.Println("Coding the Game of Life!")

	// when we run a program, the command line arguments (CLAs) go into an array called os.Args

	// we want to dictate that the length of the array is 5
	if len(os.Args) != 5 {
		panic("Incorrect number of command line parameters given. Must be 4.")
	}

	initialBoardFile := os.Args[1] // starting GameBoard file name
	outputFile := os.Args[2]       // where to draw the animated GIF

	// how many pixels wide should each cell be?
	cellWidth, err := strconv.Atoi(os.Args[3])
	if err != nil {
		panic("Error: Problem parsing cell width parameter.")
	}

	// how many generations to play the automaton?
	numGens, err2 := strconv.Atoi(os.Args[4])
	if err2 != nil {
		panic("Error: Problem parsing number of generations parameter.")
	}

	fmt.Println("Parameters read in successfully!")

	initialBoard := ReadBoardFromFile(initialBoardFile)

	fmt.Println("Playing the automaton.")

	boards := PlayGameOfLife(initialBoard, numGens)

	fmt.Println("Game of Life played. Now, drawing images.")

	// we need a slice of image objects
	imglist := DrawGameBoards(boards, cellWidth)
	fmt.Println("boards drawn to images! Convert to animated GIF.")

	// convert imagse to a GIF
	gifhelper.ImagesToGIF(imglist, outputFile)

	fmt.Println("Success! GIF produced.")

	// drawing multiple boards of the game of life
	/*
				rPentomino := ReadBoardFromFile("boards/rPentomino.csv")

				cellWidth := 20
				img := DrawGameBoard(rPentomino, cellWidth)

				filename := "output/rPentomino.png"

				f, err := os.Create(filename)
				Check(err)
				defer f.Close()

				err = png.Encode(f, img)
				Check(err)

			dinnerTable := ReadBoardFromFile("boards/dinnerTable.csv")

			cellWidth := 20
			img := DrawGameBoard(dinnerTable, cellWidth)

			filename := "output/dinnerTable.png"

			f, err := os.Create(filename)
			Check(err)
			defer f.Close()

			err = png.Encode(f, img)
			Check(err)

		gosperGun := ReadBoardFromFile("boards/gospergun.csv")

		cellWidth := 20
		img := DrawGameBoard(gosperGun, cellWidth)

		filename := "output/gospergun.png"

		f, err := os.Create(filename)
		Check(err)
		defer f.Close()

		err = png.Encode(f, img)
		Check(err)
	*/
}
