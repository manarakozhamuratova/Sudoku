package main

import (
	"fmt"
	"os"
	"sudoku"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		args = sudoku.Split(args[0])
	}
	if sudoku.IsValid(args) {
		if !sudoku.Dublicate(args) {
			table := sudoku.ConvertToInt(args)

			if sudoku.Solve(&table) {
				sudoku.ShowSudoku(table)
			} else {
				fmt.Printf("Error\n")
			}
		} else {
			fmt.Printf("Error\n")
		}
	} else {
		fmt.Printf("Error\n")
	}
}
