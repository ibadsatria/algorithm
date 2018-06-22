package bfs

type graph struct {
	vertices []*vertice
}

type vertice struct {
	visited  bool
	adjacent []*vertice
}

// NewVertice init new vertice
func NewVertice() *vertice {
	return &vertice{visited: false}
}
