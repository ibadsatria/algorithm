package bfs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddEdge(t *testing.T) {
	var g graph
	listVer := []Vertex{
		Vertex{0, false},
		Vertex{1, false},
		Vertex{2, false},
		Vertex{3, false},
	}

	t.Run("TestAddEdge", func(t *testing.T) {
		g.AddEdge(&listVer[0], &listVer[1])
		g.AddEdge(&listVer[0], &listVer[2])
		g.AddEdge(&listVer[1], &listVer[2])
		g.AddEdge(&listVer[2], &listVer[0])
		g.AddEdge(&listVer[2], &listVer[3])
		g.AddEdge(&listVer[3], &listVer[3])

		//for vertex 0
		assert.Contains(t, g.vertexEdge[&listVer[0]], &listVer[1])
		assert.Contains(t, g.vertexEdge[&listVer[0]], &listVer[2])

		// for Vertex 1
		assert.Contains(t, g.vertexEdge[&listVer[1]], &listVer[2])

		// for vertex 2
		assert.Contains(t, g.vertexEdge[&listVer[2]], &listVer[0])
		assert.Contains(t, g.vertexEdge[&listVer[2]], &listVer[3])

		// for vertex 3
		assert.Contains(t, g.vertexEdge[&listVer[3]], &listVer[3])
	})
}

func ExampleTraverse() {
	var g graph
	listVer := []Vertex{
		Vertex{0, false},
		Vertex{1, false},
		Vertex{2, false},
		Vertex{3, false},
	}

	g.AddEdge(&listVer[0], &listVer[1])
	g.AddEdge(&listVer[0], &listVer[2])
	g.AddEdge(&listVer[1], &listVer[2])
	g.AddEdge(&listVer[2], &listVer[0])
	g.AddEdge(&listVer[2], &listVer[3])
	g.AddEdge(&listVer[3], &listVer[3])

	result := g.Traverse(&listVer[2])
	fmt.Println(result)
	// Output: [2 0 3 1]
}
