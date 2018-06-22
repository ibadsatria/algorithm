package bfs

// Graph domain
type graph struct {
	vertexEdge map[Vertex][]*Vertex
}

// Traverse graph bfs manner
func (g *graph) Traverse(s *Vertex) (result []int) {
	s.Visited = true
	result = append(result, s.ID)

	for key, value := range g.vertexEdge {
		if !key.Visited {
			key.Visited = true

			for _, v := range value {
				if !v.Visited {
					v.Visited = true
					result = append(result, v.ID)
				}

			}
		}
	}
	return
}

// AddEdge add connection between two vertex
func (g *graph) AddEdge(verDomain, verRange Vertex) {
	if g.vertexEdge == nil {
		g.vertexEdge = make(map[Vertex][]*Vertex, 0)
	}

	if _, ok := g.vertexEdge[verDomain]; !ok {
		g.vertexEdge[verDomain] = []*Vertex{&verRange}
	} else {
		g.vertexEdge[verDomain] = append(g.vertexEdge[verDomain], &verRange)
	}
}
