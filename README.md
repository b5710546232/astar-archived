# A* Search in Go
This is an implementation of the A* search algorithm in Go, a popular algorithm for finding the shortest path between two points in a grid.

# Usage
To use the A* search algorithm, create a Grid type representing the search space and a Vector type representing a position in the grid. The Grid type should have a method Neighbors that returns the neighbors of a given position, and the Vector type should have a method ManhattanDistance that calculates the Manhattan distance between two vectors.

Then, create an instance of the AStar type and call the Search method to perform the search. The Search method returns a slice of vectors representing the shortest path from the start position to the end position, if it exists.

Here is an example of how to use the A* search algorithm:
```golang
// Create a grid representing the search space
grid := Grid{
	{0, 0, 0, 0, 0},
	{0, 1, 1, 1, 0},
	{0, 0, 0, 0, 0},
	{0, 0, 1, 0, 0},
	{0, 0, 0, 0, 0},
}

// Create the A* search instance
aStar := NewAStar(grid, Vector{0, 0}, Vector{4, 4})

// Perform the search
path := aStar.Search()

// Print the path, if it was found
if path != nil {
	fmt.Println("Path found:", path)
} else {
	fmt.Println("No path found")
}

```

This will search for the shortest path from the start position (0, 0) to the end position (4, 4) in the grid.

## Requirements

Go 1.15 or higher