package main

import (
	"fmt"
	"image/png" // for drawing an image to PNG
	"os"        // for opening a file that will hold the image
)

func main() {
	fmt.Println("Coding the Game of Life!")
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
	*/

	gosperGun := ReadBoardFromFile("boards/gospergun.csv")

	cellWidth := 20
	img := DrawGameBoard(gosperGun, cellWidth)

	filename := "output/gospergun.png"

	f, err := os.Create(filename)
	Check(err)
	defer f.Close()

	err = png.Encode(f, img)
	Check(err)
}
