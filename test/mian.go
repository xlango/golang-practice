package main

import (
	"fmt"
	"os"
)

func main1() {
	file, err := os.Open("test/maze.in")
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	mazes := make([][]int, row)
	for i := 0; i < len(mazes); i++ {
		mazes[i] = make([]int, col)
		for j := 0; j < len(mazes[i]); j++ {
			fmt.Fscanf(file, "%d", &mazes[i][j])
		}
	}

	//for _,maze:=range mazes{
	//	for _,m:=range maze{
	//		fmt.Printf("%d  ",m)
	//	}
	//	fmt.Println()
	//}
	pathMaze := walk(mazes, Point{0, 0}, Point{len(mazes) - 1, len(mazes[0]) - 1})

	for _, maze := range pathMaze {
		for _, m := range maze {
			fmt.Printf("%3d  ", m)
		}
		fmt.Println()
	}
}

type Point struct {
	i, j int
}

//上左下右移动
var moves = [4]Point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p Point) Move(m Point) Point {
	return Point{p.i + m.i, p.j + m.j}
}

func (p Point) checkPoint(mazes [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(mazes) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(mazes[p.i]) {
		return 0, false
	}
	return mazes[p.i][p.j], true
}

//迷宫算法
func walk(mazes [][]int, start, end Point) [][]int {
	steps := make([][]int, len(mazes))
	for i := range mazes {
		steps[i] = make([]int, len(mazes[i]))
	}
	Q := []Point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}
		for _, move := range moves {
			step := cur.Move(move)

			v, ok := step.checkPoint(mazes)
			if !ok || v == 1 {
				continue
			}

			v, ok = step.checkPoint(steps)
			if !ok || v != 0 {
				continue
			}
			if step == start {
				continue
			}

			steps[step.i][step.j] = steps[cur.i][cur.j] + 1

			Q = append(Q, step)

		}
	}

	return steps
}
