package main

import (
	"fmt"
	"os"

	exos "github.com/yprosper/hackerrank/exercices"
)

func printBoard(b [][]string) {
	for _, v := range b {
		for _, v2 := range v {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func updateBoard(x, y int, b *([][]string), action string) (int, int) {
	rx, ry := x, y
	switch action {
	case "CLEAN":
		{
			(*b)[x][y] = "b"
		}
	case "RIGHT":
		{
			(*b)[x][y] = "-"
			if (*b)[x][y+1] != "d" {
				(*b)[x][y+1] = "b"
			}
			ry++
		}
	case "LEFT":
		{
			(*b)[x][y] = "-"
			if (*b)[x][y-1] != "d" {
				(*b)[x][y-1] = "b"
			}
			ry--
		}
	case "UP":
		{
			(*b)[x][y] = "-"
			if (*b)[x-1][y] != "d" {
				(*b)[x-1][y] = "b"

			}
			rx--
		}
	case "DOWN":
		{
			(*b)[x][y] = "-"
			if (*b)[x+1][y] != "d" {
				(*b)[x+1][y] = "b"
			}
			rx++
		}
	}
	return rx, ry
}

func main() {
	exos.NPuzzle(os.Args[1])
	/*
		input, _ := os.Open("botcleanInput.txt")
		defer input.Close()

		var bx, by, h, w int
		fmt.Fscanf(input, "%d %d\n%d %d\n", &bx, &by, &h, &w)

		board := make([][]string, h)
		for i := 0; i < h; i++ {
			var s string
			fmt.Fscanf(input, "%s\n", &s)
			board[i] = strings.Split(s, "")
		}

		printBoard(board)
		i := 0
		for i < 6 {
			action := exos.Botclean(bx, by, h, w, board)
			fmt.Println(action)
			bx, by = updateBoard(bx, by, &board, action)
			printBoard(board)
			i++
		}
	*/
}
