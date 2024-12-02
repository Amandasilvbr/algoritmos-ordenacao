package main

import (
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

// Function to generate graphs
func generateGraph(algorithms []struct {
	name   string
	sorter func([]int) (int, int) // Sorting function that returns comparisons and swaps
}, sizes []int) {
	// Creating the graph
	p := plot.New()

	// Setting the title and labels for the X and Y axes
	p.Title.Text = "Comparação dos Algoritmos"
	p.X.Label.Text = "Tamanho do Array"
	p.Y.Label.Text = "Tempo (ms)"

	// Defining a vibrant and contrasting color palette
	colors := []color.Color{
		color.RGBA{R: 31, G: 119, B: 180, A: 255},  // Blue
		color.RGBA{R: 255, G: 127, B: 14, A: 255},  // Orange
		color.RGBA{R: 44, G: 160, B: 44, A: 255},   // Green
		color.RGBA{R: 214, G: 39, B: 40, A: 255},   // Red
		color.RGBA{R: 148, G: 103, B: 189, A: 255}, // Purple
		color.RGBA{R: 140, G: 86, B: 75, A: 255},   // Brown
		color.RGBA{R: 227, G: 119, B: 194, A: 255}, // Pink
		color.RGBA{R: 127, G: 127, B: 127, A: 255}, // Gray
		color.RGBA{R: 188, G: 189, B: 34, A: 255},  // Yellow
		color.RGBA{R: 23, G: 190, B: 207, A: 255},  // Cyan
	}

	// Loop through each algorithm to create a line in the graph
	for i, algo := range algorithms {

		var timePoints plotter.XYs // Stores the (X, Y) points for the graph

		// Loop over each array size and collect the execution time
		for _, size := range sizes {
			arr := generateRandomArray(size)                     // Generate a random array for the current size
			duration, _, _ := runSortAlgorithm(algo.sorter, arr) // Run the sorting algorithm and get the duration

			// Append the time (in milliseconds) for the current array size to timePoints
			timePoints = append(timePoints, plotter.XY{
				X: float64(size),
				Y: float64(duration.Milliseconds()), // Convert duration to milliseconds
			})
		}

		// Creating the line for the graph
		line, err := plotter.NewLine(timePoints)
		if err != nil {
			fmt.Println("Error creating line:", err)
			return
		}

		// Applying the line color
		line.Color = colors[i]
		line.Width = 2

		// Adding the line to the graph
		p.Add(line)
		p.Legend.Add(algo.name, line)
	}

	// Saving the graph as a PNG file
	if err := p.Save(10*vg.Inch, 6*vg.Inch, "sorting_algorithms_time.png"); err != nil {
		fmt.Println("Error saving plot:", err)
	}
}
