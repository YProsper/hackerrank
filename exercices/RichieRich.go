package exercices

import (
	"fmt"
	"os"
)

// RichieRich bla bla bla
func RichieRich(inputFile, outputFile string) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	output, _ := os.Open(outputFile)
	defer output.Close()

	var T int
	fmt.Fscanf(input, "%d\n", &T)

	for i := 0; i < T; i++ {
		var n, k int
		fmt.Fscanf(input, "%d %d\n", &n, &k)

		var s string
		fmt.Fscanf(input, "%s\n", &s)

		var expected string
		fmt.Fscanf(output, "%s\n", &expected)

		l, r, c := 0, len(s)-1, 0
		a := make([]byte, len(s))
		copy(a, s)

		for l <= r {
			if s[l] > s[r] {
				a[l] = s[l]
				a[r] = s[l]
				c++
			} else if s[l] < s[r] {
				a[l] = s[r]
				a[r] = s[r]
				c++
			}
			l++
			r--
		}

		if c > k {
			fmt.Println(-1)
			break
		} else if k-c == 1 && len(a)%2 != 0 && a[len(a)/2] != '9' {
			a[len(s)/2] = '9'
		} else if c < k {
			i, j := 0, len(s)-1
			for i <= j && c < k {
				if a[i] != s[i] && a[i] != '9' {
					a[i] = '9'
					a[j] = '9'
					c++
				} else if a[j] != s[j] && a[j] != '9' {
					a[i] = '9'
					a[j] = '9'
					c++
				} else if a[i] != '9' {
					a[i] = '9'
					a[j] = '9'
					c += 2
				}

				i++
				j--
			}
		}

		fmt.Println(string(a))
	}
}
