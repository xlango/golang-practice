package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	//打开文件
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int,
	start, end point) [][]int {
	//保存已走位置
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	//队列保存  可向下一步探索的当前站立点
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		//分别向上、左、下、右探索
		for _, dir := range dirs {
			//探索到的位置
			next := cur.add(dir)

			//判断探索到的位置 是否为撞墙 或者 是否越界
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] =
				curSteps + 1

			//将下一批可站立并向前探索的点放入队列
			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze("labyrinth/maze.in")
}

func maze(filename string) {
	//读取文件中的迷宫
	maze := readMaze(filename)

	//steps := walk(maze, point{0, 0},
	//	point{len(maze) - 1, len(maze[0]) - 1})
	steps := walk(maze, point{0, 4},
		point{0, 0})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
