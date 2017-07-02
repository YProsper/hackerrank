package main

import (
	"bufio"
	"fmt"
	"os"
)

// AbsolutePermutation bla bla bla
func AbsolutePermutation(inputFile, outputFile string) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	output, _ := os.Open(outputFile)
	defer output.Close()

	scannerOutput := bufio.NewScanner(output)
	scannerOutput.Split(bufio.ScanLines)

	var T int
	fmt.Fscanf(input, "%d\n", &T)

	for i := 0; i < T; i++ {
		var n, k int
		fmt.Fscanf(input, "%d %d\n", &n, &k)
		fmt.Printf("%d %d\n", n, k)

		scannerOutput.Scan()
		expected := scannerOutput.Text()
		fmt.Printf("expected=%s\n", expected)

		a := make([]int, n+1)
		b := make([]bool, n+1)
		finish := false

		j := 1
		for j < n+1 && !finish {
			x1 := (j - k)
			if x1 < 0 {
				x1 += n
			}

			x2 := (j + k)
			if x2 > n {
				x2 -= n
			}

			if x1 == 0 || b[x1] {
				if x2 == 0 || b[x2] {
					fmt.Println(-1)
					finish = true
				} else {
					a[j] = x2
					b[x2] = true
				}
			} else if x2 == 0 || b[x2] {
				a[j] = x1
				b[x1] = true
			} else if x1 < x2 {
				a[j] = x1
				b[x1] = true
			} else {
				a[j] = x2
				b[x2] = true
			}

			j++
		}

		if !finish {
			for i := 1; i < n+1; i++ {
				fmt.Printf("%d ", a[i])
			}
			fmt.Println("")
		}
	}
}

func main() {
	AbsolutePermutation(os.Args[1], os.Args[2])
}
