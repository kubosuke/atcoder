package main

import (
	"fmt"
)

type Point struct {
	Y int
	X int
}

var h, w int

func main() {
	fmt.Scanf("%d %d", &h, &w)

	var (
		tmp string
		level,maxlevel int
	)

	s := make([][]string, h)
	for i := 0; i < h; i++ {
		fmt.Scan(&tmp)
		for _, c := range tmp {
			s[i] = append(s[i], string(c))
		}
	}

	for i:=0; i<h; i++ {
		for j:=0; j<w; j++ {
			level = bfs(Point{i,j}, s)
			if maxlevel < level {
				maxlevel = level
			}
		}
	}

	fmt.Println(maxlevel)
}

func bfs(start Point, m [][]string) int {
	if m[start.Y][start.X] == "#" {
		return 0
	}
	var (
		p  Point
		queue []Point
		level int
	)
	maze := make([][]string, len(m))
	for i := range m {
		maze[i] = make([]string, len(m[i]))
		copy(maze[i], m[i])
	}
	queue = append(queue, start)
	maze[start.Y][start.X] = "#"

	for level = 0; len(queue) != 0; level++ {
		taskqueue := make([]Point, len(queue))
		copy(taskqueue, queue)
		queue = []Point{}
		for len(taskqueue) != 0 {
			p = taskqueue[:1][0]
			taskqueue = taskqueue[1:]

			// up
			if p.Y != 0 {
				if maze[p.Y-1][p.X] != "#" {
					queue = append(queue, Point{p.Y - 1, p.X})
					maze[p.Y-1][p.X] = "#"
				}
			}

			// right
			if p.X != w-1 {
				if maze[p.Y][p.X+1] != "#" {
					queue = append(queue, Point{p.Y, p.X + 1})
					maze[p.Y][p.X+1] = "#"
				}
			}

			//left
			if p.X != 0 {
				if maze[p.Y][p.X-1] != "#" {
					queue = append(queue, Point{p.Y, p.X - 1})
					maze[p.Y][p.X-1] = "#"
				}
			}

			//down
			if p.Y != h-1 {
				if maze[p.Y+1][p.X] != "#" {
					queue = append(queue, Point{p.Y + 1, p.X})
					maze[p.Y+1][p.X] = "#"
				}
			}
		}
	}
	return level - 1
}
