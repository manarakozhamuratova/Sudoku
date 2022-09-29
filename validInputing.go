package sudoku

func IsValid(table []string) bool {
	if len(table) != 9 {
		return false
	} else {
		for row := 0; row < len(table); row++ {
			if len(table[row]) != 9 {
				return false
			}
		}
		for row := 0; row < len(table); row++ {
			for column := 0; column < len(table); column++ {
				if !IsNum(rune(table[row][column])) && !isDot(rune(table[row][column])) {
					return false
				}
			}
		}
	}
	return true
}

func IsNum(char rune) bool {
	return char >= '1' && char <= '9'
}

func isDot(char rune) bool {
	return char == '.'
}

func Dublicate(table []string) bool {
	for k := 0; k < 9; k++ {
		for i := 0; i < 9; i++ {
			if IsNum(rune(table[k][i])) {
				for j := 0; j < 9; j++ {
					if table[k][i] == table[k][j] && i != j {
						return true
					}
				}
			}
		}
	}

	for k := 0; k < 9; k++ {
		for i := 0; i < 9; i++ {
			if IsNum(rune(table[i][k])) {
				for j := 0; j < 9; j++ {
					if table[i][k] == table[j][k] && i != j {
						return true
					}
				}
			}
		}
	}
	return false
}
