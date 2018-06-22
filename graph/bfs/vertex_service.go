package bfs

// VertexService service for vertex matter
type VertexService struct{ Vertex }

// CreateVertex return new initialized vertex
func (vs *VertexService) CreateVertex(id int) *Vertex {
	return &Vertex{ID: id, Visited: false}
}
