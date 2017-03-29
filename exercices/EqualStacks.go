package exercices

import (
	"bufio"
	"fmt"
	"os"
)

// EqualStacks is a function
func EqualStacks(inputFile, outputFile string) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	output, _ := os.Open(outputFile)
	defer output.Close()

	var n1, n2, n3 int
	fmt.Fscan(input, &n1, &n2, &n3)

	var expected int
	fmt.Fscan(output, &expected)

	s1, s2, s3 := make([]int, n1), make([]int, n2), make([]int, n3)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	sum1, sum2, sum3 := 0, 0, 0

	for i := 0; i < n1; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &s1[i])
		sum1 += s1[i]
	}

	for i := 0; i < n2; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &s2[i])
		sum2 += s2[i]
	}

	for i := 0; i < n3; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &s3[i])
		sum3 += s3[i]
	}

	t1, t2, t3, result := 0, 0, 0, 0
	found := false

	for t1 < n1 && t2 < n2 && t3 < n3 && !found {
		if sum1 == sum2 && sum1 == sum3 {
			result = sum1
			found = true
		} else if sum1 >= sum2 && sum1 >= sum3 {
			sum1 -= s1[t1]
			t1++
		} else if sum2 >= sum1 && sum2 >= sum3 {
			sum2 -= s2[t2]
			t2++
		} else if sum3 >= sum1 && sum3 >= sum2 {
			sum3 -= s3[t3]
			t3++
		}
	}

	fmt.Printf("%d %d\n", result, expected)
}
