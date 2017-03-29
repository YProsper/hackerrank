package exercices

import (
	"fmt"
	"math/big"
)

// FibonacciModified Tn+2 = (Tn+1)2 + Tn
func FibonacciModified(a, b, n int) string {
	t1, t2 := big.NewInt(int64(a)), big.NewInt(int64(b))
	for i := 3; i <= n; i++ {
		fmt.Printf("T%d -> %s %s -> ", i, t1.String(), t2.String())
		t3 := big.NewInt(0)
		t3.Mul(t2, t2)
		t3.Add(t3, t1)
		t1, t2 = t2, t3
		fmt.Printf("%s %s\n", t1.String(), t2.String())
	}

	return t2.String()
}
