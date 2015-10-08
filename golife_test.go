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
		neighbours cellMax
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

func TestCountNeighbours_GivenAnEmptyWorld_ReturnsZero(t *testing.T) {
	w := world{width: 5, height: 5}
	var want cellMax

	c := w.countLiveNeighbours(0)

	if c != want {
		t.Errorf("countNeighbours() == %d, want %d", c, want)
	}
}

func TestCountNeighbours_GivenAFullWorld_ReturnsCorrectIndex(t *testing.T) {
	w := world{width: 3, height: 3}
	w.cells = append(w.cells, cell{0, 0, true}, cell{0, 1, true}, cell{0, 2, true}, cell{1, 0, true}, cell{1, 1, true}, cell{1, 2, true}, cell{2, 0, true}, cell{2, 1, true}, cell{2, 2, true})
	cases := []struct {
		index, want cellMax
	}{
		{0, 3},
		{1, 5},
		{2, 3},
		{3, 5},
		{4, 8},
		{5, 5},
		{6, 3},
		{7, 5},
		{8, 3},
	}

	for _, c := range cases {
		n := w.countLiveNeighbours(c.index)

		if n != c.want {
			t.Errorf("countNeighbours(%d) == %d, want %d", c.index, n, c.want)
		}
	}
}
