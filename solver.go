package sudoku

import (
	"fmt"
)

func Solve(table *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if table[row][col] == 0 {
				for number := 1; number <= 9; number++ {
					if IsValidFilling(table, number, row, col) {
						table[row][col] = number

						if Solve(table) {
							return true
						}
						table[row][col] = 0

					}
				}

				return false
			}
		}
	}

	return true
}

func ShowSudoku(table [9][9]int) {
	for row := 0; row < len(table); row++ {
		for col := 0; col < len(table[row]); col++ {
			fmt.Print(table[row][col])
			if col != 8 {
				fmt.Print(" ")
			}

		}
		fmt.Println()
	}
	fmt.Println()
}

func ConvertToInt(table []string) [9][9]int {
	convert := [9][9]int{}
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			convert[i][j] = CharToInt(rune(table[i][j]))
		}
	}
	return convert
}

func CharToInt(char rune) int {
	if char >= '1' && char <= '9' {
		return int(char) - '0'
	} else {
		return 0
	}
}

func Split(s string) []string {
	var answer string
	var res []string
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if !IsNum(runes[i]) && !isDot(runes[i]) {
			if answer != "" {
				res = append(res, answer)
			}

			answer = ""
		} else {
			answer += string(runes[i])
		}
	}
	res = append(res, answer)
	return res
}
