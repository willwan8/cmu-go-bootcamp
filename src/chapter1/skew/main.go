package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	/* installing this was very problematic, so skipped the diagram stuff
	"pkg/mod/gonum.org/v1/plot"
	"pkg/mod/gonum.org/v1/plot/plotter" // include line graph support
	"pkg/mod/gonum.org/v1/plot/vg" // support for vector graphics and image generation
	*/
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

	/*
	// draw the skew diagram
	MakeSkewDiagram(EcoliSkewArray)
	*/
}

// MakeSkewDiagram
// Input: A skew array.
// Output: (none)
// Draws the skew diagram of the given skew array to an image and saves to file.
/* 
installing w/ go get provided too many problems, so just skipped this image forming part
func MakeSkewDiagram(skewArray []int) {
	p := plot.New() // creates new plotter object

	p.Title.Text = "Skew Diagram"
	p.X.Label.Text = "Genome Position"
	p.Y.Label.Text = "Skew value"

	// remove legend
	p.Legend.Top = false

	// make a collection of points associated with each skew value
	points := make(plotter.XYs, len(skewArray))

	// set the X and Y value of each point
	for i, val := range SkewArray {
		points[i].X = float64(i)
		points[i].Y = float64(skewValue)
	}

	// connect the dots!
	line, err := plotter.NewLine(points)

	if err != nil {
		panic(err)
	}

	// add our line to the plot
	p.Add(line)

	// draw to an image

	// first, set a unit of length
	unitOflength := vg.Centimeter

	// make label fonts bigger
	p.X.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Y.Label.TextStyle.Font.Size = 3 * unitOfLength
	p.Title.TextStyle.Font.Size = 4 * unitOfLength
	p.X.Tick.Label.Font.Size = 2 * unitOfLength
	p.Y.Tick.Label.Font.Size = 2 * unitOfLength

	// save my plot to a PNG
	err = p.Save(100*unitOflength, 60*unitOflength, "skewDiagram.png")

	if err != nil {
		panic(err)
	}
}
*/