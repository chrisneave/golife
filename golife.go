package main

type world struct {
	cells         []cell
	width, height int
}

type cell struct {
	x, y  int
	alive bool
}

func (c *cell) isAlive() (isAlive bool) {
	return c.alive
}

func (c *cell) isDead() (isAlive bool) {
	return !c.alive
}

func shouldCellLive(alive bool, neighbours int) (live bool) {
	if neighbours == 2 && alive == true {
		return true
	}
	if neighbours == 3 {
		return true
	}
	return false
}

func (w *world) countLiveNeighbours(cellIndex int) (neighbours int) {
	if cellIndex >= len(w.cells) {
		return
	}

	cell := w.cells[cellIndex]

	for i, c := range w.cells {
		// Don't count the target cell.
		if cellIndex == i {
			continue
		}

		// Only count live cells.
		if c.isDead() {
			continue
		}

		// Ignore cells that are not next to the target cell on the x axis.
		if c.x < cell.x-1 || c.x > cell.x+1 {
			continue
		}

		// Ignore cells that are not next to the target cell on the y axis.
		if c.y < cell.y-1 || c.y > cell.y+1 {
			continue
		}

		neighbours++
	}

	return
}
