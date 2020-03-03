package theknightstour

import "fmt"

// N ...
var N int

type move struct {
	x int
	y int
}

var moves []move = []move{
	{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1},
}

var solved [][]int

// Solution https://www.geeksforgeeks.org/the-knights-tour-problem-backtracking-1/
func Solution(n int) {
	N = n
	solved = make([][]int, N)
	for i := range solved {
		solved[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			bt(1, move{i, j})
		}
	}
	for _, v := range solved {
		for _, w := range v {
			fmt.Printf("%2v ", w)
		}
		fmt.Println()
	}
}

func bt(step int, curM move) bool {
	if !isSafe(curM) {
		return false
	}
	solved[curM.x][curM.y] = step
	if step == N*N {
		return true
	}
	for _, m := range moves {
		nextM := move{
			m.x + curM.x,
			m.y + curM.y,
		}
		if bt(step+1, nextM) {
			return true
		}
	}
	solved[curM.x][curM.y] = 0
	return false
}

func isSafe(m move) bool {
	if m.x < 0 || m.x >= N || m.y < 0 || m.y >= N || solved[m.x][m.y] != 0 {
		return false
	}
	return true
}
