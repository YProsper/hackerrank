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
				a[y*C+x] = 3
			}
		}
	}
}

func updateBombs(a []int) {
	for i, v := range a {
		if v > 0 {
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
				b.WriteRune('0')
				//b.WriteString(strconv.Itoa(a[y*C+x]))
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
				a[i*C+j] = 3
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
	//printBombs(a, R, C)
	//fmt.Printf("Second 1: nothing\n")
	updateBombs(a)
	//printBombs(a, R, C)
	if N%2 == 0 {
		fillBombs(a, R, C)
		printBombs(a, R, C)
		return
	}

	for i := 2; i <= N; i++ {
		updateBombs(a)
		switch i % 2 {
		case 1:
			{
				//			fmt.Printf("Second %d,%d: detonate\n", i, i%2)
				detonateBombs(a, R, C)
				//		printBombs(a, R, C)
			}
		case 0:
			{
				//			fmt.Printf("Second %d,%d: plants\n", i, i%2)
				fillBombs(a, R, C)
				//		printBombs(a, R, C)
			}
		}

		fmt.Printf("Second %d ---------\n", i)
		printBombs(a, R, C)
	}

	s1 := toString(a, R, C)
	s2 := toString(expected, R, C)

	//printBombs(a, R, C)
	fmt.Println(s1 == s2)
	//printBombs(expected, R, C)
}
