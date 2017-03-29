package exercices

import (
	"fmt"
)

func nbDiv(n int) int {
	m := make(map[int]int)

	for i := 2; i <= n; i++ {
		for (n % i) == 0 {
			m[i]++
			n /= i
		}
	}

	result := 1
	for _, v := range m {
		result *= (v + 1)
	}

	return result
}

// AntiprimeNumber...
func AntiprimeNumber(n int, ints []int) {
	champion, p := 0, 0
	a := make([]int, 10000)
	for i := 0; i < 100000000; i++ {
		t := nbDiv(i)
		if t > champion {
			champion = t
			a[p] = i
			p++
			fmt.Printf("We found an antiprime : %d\n", i)
		}
	}

	fmt.Println(a)
	for i := 0; i < n; i++ {
		j := 0
		for a[j] < ints[i] {
			j++
		}
		fmt.Println(a[j])
	}
}
