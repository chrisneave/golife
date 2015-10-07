package main

func shouldCellLive(alive bool, neighbours int) (live bool) {
	if neighbours == 2 && alive == true {
		return true
	}
	if neighbours == 3 {
		return true
	}
	return false
}
