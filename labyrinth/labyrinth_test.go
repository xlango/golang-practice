package main

import (
	"testing"
)

func BenchmarkMaze(b *testing.B) {
	maze("labyrinth/maze.in")
}
