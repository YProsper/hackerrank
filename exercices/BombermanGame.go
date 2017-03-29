package exercices

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func fillBombs(a []int, R, C int) {
	for y := 0; y < R; y++ {
		for x := 0; x < C; x++ {
			if a[y*C+x] == -1 {
				a[y*C+x] = 2
			}
		}
	}
}

func updateBombs(a []int) {
	for i, v := range a {
		if v != -1 && v > 0 {
			a[i]--
		}
	}
}

func detonateBombs(a []int, R, C int) {
	for y := 0; y < R; y++ {
		for x := 0; x < C; x++ {
			if a[y*C+x] == 0 {
				a[y*C+x] = -1

				if x-1 >= 0 && a[y*C+(x-1)] != 0 {
					a[y*C+(x-1)] = -1
				}

				if y-1 >= 0 && a[(y-1)*C+x] != 0 {
					a[(y-1)*C+x] = -1
				}

				if x+1 < R && a[y*C+(x+1)] != 0 {
					a[y*C+(x+1)] = -1
				}

				if y+1 < R && a[(y+1)*C+x] != 0 {
					a[(y+1)*C+x] = -1
				}
			}
		}
	}
}

func printBombs(a []int, R, C int) {
	s := toString(a, R, C)
	fmt.Print(s)
}

func toString(a []int, R, C int) string {
	var b bytes.Buffer
	for y := 0; y < R; y++ {
		for x := 0; x < C; x++ {
			if a[y*C+x] == -1 {
				b.WriteRune('.')
			} else {
				b.WriteRune('O')
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}

func printBombsDebug(a [][]int, R, C int) {
	for y := 0; y < R; y++ {
		for x := 0; x < C; x++ {
			fmt.Print(a[y*C+x])
		}
		fmt.Println("")
	}
	fmt.Println("")
}

func readBombs(scanner *bufio.Scanner, R, C int) []int {
	a := make([]int, R*C)
	for i := 0; i < R; i++ {
		scanner.Scan()
		for j, v := range scanner.Text() {
			switch v {
			case '.':
				a[i*C+j] = -1
			case 'O':
				a[i*C+j] = 2
			}
		}
	}

	return a
}

// BombermanGame bla bla bla
func BombermanGame(inputFile, outputFile string) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	output, _ := os.Open(outputFile)
	defer output.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	scannerOutput := bufio.NewScanner(output)
	scannerOutput.Split(bufio.ScanWords)

	var R, C, N int
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &R)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &C)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "%d", &N)

	a := readBombs(scanner, R, C)
	expected := readBombs(scannerOutput, R, C)

	updateBombs(a)
	for i := 1; i <= N; i++ {
		switch i % 2 {
		case 0:
			{
				detonateBombs(a, R, C)
			}
		case 1:
			{
				fillBombs(a, R, C)
			}
		}
		updateBombs(a)
	}

	s1 := toString(a, R, C)
	s2 := toString(expected, R, C)
	fmt.Println(s1 == s2)
	//printBombs(a, R, C)
	//printBombs(expected, R, C)
}
