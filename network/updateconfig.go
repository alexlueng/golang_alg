package network

import "container/list"

// feature 9 Determine the time for an update request to propagate through all routers in a grid topology

type configPair struct {
	first, second int
}

func updateCOnfiguration(grid [][]int) int {
	queue := list.New()

	rows, cols := len(grid), len(grid[0])
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue.PushBack(Pair{first: r, second: c})
			}
		}
	}

	queue.PushBack(Pair{first: -1, second: -1})

	minutesElapsed := -1

	directions := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for queue.Len() != 0 {
		p := queue.Front()
		queue.Remove(p)
		q := p.Value.(Pair)
		row := q.first
		col := q.second

		if row == -1 {
			minutesElapsed++
			if queue.Len() != 0 {
				queue.PushBack(Pair{first: -1, second: -1})
			}
		} else {
			for _, d := range directions {
				neighborRow := row + d[0]
				neighborCol := col + d[0]
				if neighborRow >= 0 && neighborRow < rows && neighborCol >= 0 && neighborCol < cols {
					if grid[neighborRow][neighborCol] == 1 {
						grid[neighborRow][neighborCol] = 2
						queue.PushBack(Pair{first: neighborRow, second: neighborCol})
					}
				}
			}
		}
	}
	return minutesElapsed
}
