package bfs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewVertex(t *testing.T) {
	tests := map[int]Vertex{
		1:   Vertex{1, false},
		0:   Vertex{0, false},
		4:   Vertex{4, false},
		-99: Vertex{0, false},
	}

	for k, v := range tests {
		t.Run("TestNewVertex", func(t *testing.T) {
			verPointer := CreateVertex(k)
			assert.Equal(t, v.ID, verPointer.ID)
			assert.Equal(t, v.Visited, verPointer.Visited)
		})
	}
}
