package bfs

// GraphService service for graph matter
type GraphService struct {
	VertexService
	Graph
}

// Traverse graph bfs manner
func (gs *GraphService) Traverse() {

}

// AddEdge add connection between two vertex
func (gs *GraphService) AddEdge(verDomain, verRange Vertex) {
	if _, ok := gs.VertexEdge[verDomain]; !ok {
		gs.VertexEdge[verDomain] = []*Vertex{&verRange}
	} else {
		gs.VertexEdge[verDomain] = append(gs.VertexEdge[verDomain], &verRange)
	}
}
