package astar

// Vector represents a position in the grid
type Vector struct {
	X int
	Y int
}

// ManhattanDistance calculates the Manhattan distance between two vectors
func (v Vector) ManhattanDistance(other Vector) int {
	return abs(v.X-other.X) + abs(v.Y-other.Y)
}
func (v Vector) Add(other Vector) Vector {
	return Vector{v.X + other.X, v.Y + other.Y}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Node represents a node in the search graph
type Node struct {
	// The position of the node
	Pos Vector

	// The cost of getting from the start node to this node
	G int

	// The estimated cost of getting from this node to the end node
	H int

	// The total cost of getting from the start node to the end node through this node
	F int

	// The parent node of this node, used to reconstruct the path
	Parent *Node
}

// AStar implements the A* search algorithm
type AStar struct {
	// The grid representing the search space
	Grid Grid

	// The start and end nodes
	Start, End *Node

	// The set of open nodes
	OpenSet map[Vector]*Node

	// The set of closed nodes
	ClosedSet map[Vector]*Node
}

// Grid represents the search space
type Grid [][]int

// Neighbors returns the neighbors of a given position
func (g Grid) Neighbors(pos Vector) []Vector {
	var neighbors []Vector
	for _, dir := range dirs {
		neighbor := pos.Add(dir)
		if g.IsValidPos(neighbor) && g[neighbor.Y][neighbor.X] != 1 {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

// IsValidPos returns true if the given position is within the bounds of the grid
func (g Grid) IsValidPos(pos Vector) bool {
	return pos.X >= 0 && pos.X < len(g[0]) && pos.Y >= 0 && pos.Y < len(g)
}

var dirs = []Vector{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

// NewAStar creates a new A* search instance
func NewAStar(grid Grid, start, end Vector) *AStar {
	startNode := &Node{Pos: start, G: 0, H: start.ManhattanDistance(end), F: start.ManhattanDistance(end)}
	endNode := &Node{Pos: end}
	return &AStar{
		Grid:      grid,
		Start:     startNode,
		End:       endNode,
		OpenSet:   map[Vector]*Node{start: startNode},
		ClosedSet: map[Vector]*Node{},
	}
}

// Search performs the A* search and returns the path from start to end, if it exists
func (a *AStar) Search() []Vector {
	for len(a.OpenSet) > 0 {
		// Get the node in the open set with the lowest F score
		current := a.lowestFScore()

		// If we have reached the end, reconstruct the path and return it
		if current.Pos.X == a.End.Pos.X && current.Pos.Y == a.End.Pos.Y {
			return a.reconstructPath(current)
		}

		// Remove the current node from the open set and add it to the closed set
		delete(a.OpenSet, current.Pos)
		a.ClosedSet[current.Pos] = current

		// Get the neighbors of the current node
		neighbors := a.Grid.Neighbors(current.Pos)

		// Iterate through the neighbors
		for _, neighborPos := range neighbors {
			// If the neighbor is already in the closed set, skip it
			if _, ok := a.ClosedSet[neighborPos]; ok {
				continue
			}

			// Calculate the cost of moving from the current node to the neighbor
			cost := current.G + current.Pos.ManhattanDistance(neighborPos)

			// Get the neighbor node from the open set, if it exists
			neighbor, ok := a.OpenSet[neighborPos]

			// If the neighbor is not in the open set, or if the cost of reaching it from the current node
			// is lower than the previously calculated cost, update its cost and set its parent to the current node
			if !ok || cost < neighbor.G {
				neighbor = &Node{Pos: neighborPos, G: cost, H: neighborPos.ManhattanDistance(a.End.Pos), F: cost + neighborPos.ManhattanDistance(a.End.Pos), Parent: current}
				a.OpenSet[neighborPos] = neighbor
			}
		}
	}

	// If we get here, the search has failed and there is no path from start to end
	return nil
}

// lowestFScore returns the node in the open set with the lowest F score
func (a *AStar) lowestFScore() *Node {
	var lowest *Node
	for _, node := range a.OpenSet {
		if lowest == nil || node.F < lowest.F {
			lowest = node
		}
	}
	return lowest
}

// reconstructPath reconstructs the path from the end node back to the start node
func (a *AStar) reconstructPath(end *Node) []Vector {
	path := []Vector{end.Pos}
	current := end
	for current.Parent != nil {
		path = append(path, current.Parent.Pos)
		current = current.Parent
	}
	// Reverse the path so that it starts at the start node and ends at the end node
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
