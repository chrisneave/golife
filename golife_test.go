package main

import "testing"

/*
1. Any live cell with fewer than two live neighbours dies, as if caused by under-population.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by over-population.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
*/

func TestShouldCellLive(t *testing.T) {
	cases := []struct {
		alive      bool
		neighbours int
		live       bool
	}{
		{false, 0, false},
		{true, 0, false},
		{false, 1, false},
		{true, 1, false},
		{false, 2, false},
		{true, 2, true},
		{false, 3, true},
		{true, 3, true},
		{false, 4, false},
		{true, 4, false},
		{false, 5, false},
		{true, 5, false},
		{false, 6, false},
		{true, 6, false},
		{false, 7, false},
		{true, 7, false},
		{false, 8, false},
		{true, 8, false},
		{false, 9, false},
		{true, 9, false},
	}

	for _, c := range cases {
		live := shouldCellLive(c.alive, c.neighbours)
		if live != c.live {
			t.Errorf("shouldCellLive(%t, %d) == %t, want %t", c.alive, c.neighbours, live, c.live)
		}
	}
}
