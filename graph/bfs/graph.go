package bfs

// Graph domain
type graph struct {
	vertexEdge map[*Vertex][]*Vertex
}

// Traverse graph bfs manner
// BFS algorithm
func (g *graph) Traverse(s *Vertex) (result []int) {
	var queue []*Vertex
	s.Visited = true
	queue = append(queue, s)

	for len(queue) != 0 {
		temp := queue[:1][0]
		queue = queue[1:]
		result = append(result, temp.ID)
		for _, v := range g.vertexEdge[temp] {
			if !v.Visited {
				v.Visited = true
				queue = append(queue, v)
			}
		}
	}

	return
}

// AddEdge add connection between two vertex
func (g *graph) AddEdge(verDomain, verRange *Vertex) {
	if g.vertexEdge == nil {
		g.vertexEdge = make(map[*Vertex][]*Vertex, 0)
	}

	if _, ok := g.vertexEdge[verDomain]; !ok {
		g.vertexEdge[verDomain] = []*Vertex{verRange}
	} else {
		g.vertexEdge[verDomain] = append(g.vertexEdge[verDomain], verRange)
	}
}
