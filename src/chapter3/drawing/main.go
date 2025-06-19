package main

import (
	"chapter3/canvas"
	"fmt"
)

func main() {
	fmt.Println("Drawing a head.")

	DrawFace()
}

// DrawFace takes no inputs and returns nothing.
// It uses the canvas package to draw a face to a file
// in the output folder.
func DrawFace() {
	// create the canvas
	c := canvas.CreateNewCanvas(1000, 2000)

	// fill canvas as black
	black := canvas.MakeColor(0, 0, 0)
	c.SetFillColor(black)
	c.Clear()

	white := canvas.MakeColor(255, 255, 255)
	c.SetFillColor(white)

	// make face
	c.Circle(500, 500, 200) // x-coord 500, y-coord 500 (y goes downwards), radius 200
	c.Fill()

	// make nose
	c.SetFillColor(black)
	c.Circle(500, 550, 10)
	c.Fill()

	// make eyes
	c.Circle(425, 475, 15)
	c.Circle(575, 475, 15)
	c.Fill()

	// make mouth
	red := canvas.MakeColor(255, 0, 0)
	c.SetFillColor(red)
	c.ClearRect(400, 600, 600, 620) // x-coords range from 400-600, y-coords range from 600-620

	// save the image associated with the canvas
	c.SaveToPNG("fun.png")
}
