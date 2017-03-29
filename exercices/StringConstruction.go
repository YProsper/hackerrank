package exercices

import (
	"fmt"
	"os"
)

// StringConstruction bla bla bla
func StringConstruction(inputFile, outputFile string) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	output, _ := os.Open(outputFile)
	defer output.Close()

	var n int
	fmt.Fscanf(input, "%d\n", &n)

	for i := 0; i < n; i++ {
		var st string
		fmt.Fscanln(input, &st)

		fmt.Println(len(st))

		var st2 string
		fmt.Fscanln(output, &st2)

		a := make([]bool, 26)

		for _, v := range st {
			a[v-'a'] = true
		}
		result := 0
		for _, v := range a {
			if v {
				result++
			}
		}

		fmt.Printf("%d %s\n", result, st2)
	}
}
