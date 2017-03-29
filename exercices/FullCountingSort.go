package exercices

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// FullCountingSort bla bla bla
func FullCountingSort(inputFile string) {
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	var n int
	fmt.Sscan(scanner.Text(), &n)
	b := make([]int, n)
	c := make([]string, n)

	t0 := time.Now()
	for i := 0; i < n; i++ {
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%d", &b[i])
		scanner.Scan()
		fmt.Sscanf(scanner.Text(), "%s", &c[i])
	}

	t1 := time.Now()
	fmt.Printf("The call took %v to read file.\n", t1.Sub(t0))

	t0 = time.Now()
	helper := make([]int, 100)
	for _, v := range b {
		helper[v]++
	}
	for i := 1; i < 100; i++ {
		helper[i] += helper[i-1]
	}
	t1 = time.Now()
	fmt.Printf("The call took %v to build helper.\n", t1.Sub(t0))

	t0 = time.Now()
	r := make([]string, n)
	j, t := len(b)-1, 0
	for j >= n/2 {
		t = b[j]
		r[helper[t]-1] = c[j]
		helper[t]--
		j--
	}
	for j >= 0 {
		t = b[j]
		r[helper[t]-1] = "-"
		helper[t]--
		j--
	}
	t1 = time.Now()
	fmt.Printf("The call took %v to build r.\n", t1.Sub(t0))
	//fmt.Print(strings.Join(r, " "))
}
