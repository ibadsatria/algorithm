package bfs

// Vertex domain
type Vertex struct {
	ID      int
	Visited bool
}

// CreateVertex return new initialized vertex
// id must be greater and equal to 1
func CreateVertex(id int) *Vertex {
	if id < 1 {
		return &Vertex{ID: 0, Visited: false}
	}
	return &Vertex{ID: id, Visited: false}
}
