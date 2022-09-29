package sudoku

func IsInBox(table *[9][9]int, num, horizontal, vertical int) bool {
	rowStart := horizontal - horizontal%3
	colStart := vertical - vertical%3
	for row := rowStart; row < rowStart+3; row++ {
		for col := colStart; col < colStart+3; col++ {
			if table[row][col] == num {
				return true
			}
		}
	}
	return false
}

func IsInRow(table *[9][9]int, num, vertical int) bool {
	for col := 0; col < 9; col++ {
		if table[vertical][col] == num {
			return true
		}
	}
	return false
}

func IsInCol(table *[9][9]int, num, horizontal int) bool {
	for row := 0; row < 9; row++ {
		if table[row][horizontal] == num {
			return true
		}
	}
	return false
}

func IsValidFilling(table *[9][9]int, num, horizontal, vertical int) bool {
	return !IsInCol(table, num, vertical) && !IsInRow(table, num, horizontal) && !IsInBox(table, num, horizontal, vertical)
}
