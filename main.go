package main

import (
	"astar/astar"
	"fmt"
)

func main() {
	/**
	This example creates a 5x5 grid representing the search space, with walls represented by the value 1.
	It then creates an instance of the A* search algorithm, passing in the grid and the start and end positions.
	Finally, it calls the Search method to perform the search and prints the path if it was found.
	**/
	// Create a grid representing the search space
	grid := astar.Grid{
		{0, 0, 0, 0, 0},
		{0, 1, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}

	// Create the A* search instance
	aStar := astar.NewAStar(grid, astar.Vector{0, 0}, astar.Vector{4, 4})

	// Perform the search
	path := aStar.Search()

	// Print the path, if it was found
	if path != nil {
		fmt.Println("Path found:", path)
	} else {
		fmt.Println("No path found")
	}
}
