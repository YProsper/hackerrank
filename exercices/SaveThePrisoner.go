package exercices

// SaveThePrisoner blablabal
func SaveThePrisoner(n, m, s int) int {
	r := s - 1 + m
	for r > n {
		r -= n
	}

	return r
}
