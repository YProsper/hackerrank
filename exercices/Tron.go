package exercices

import (
	"fmt"
	"os"
	"strings"
)

type state struct {
	xr, yr, xg, yg int
	m              [][]string
}

func (s *state) isValid(x, y int) bool {
	notOutside := (x < 15 && x >= 0 && y < 15 && y >= 0)
	return notOutside && s.m[x][y] == "-"
}

func (s *state) print() {
	for i := range s.m {
		fmt.Println(s.m[i])
	}
}

func up(x, y int) (int, int) {
	return x - 1, y
}

func down(x, y int) (int, int) {
	return x + 1, y
}

func left(x, y int) (int, int) {
	return x, y - 1
}

func right(x, y int) (int, int) {
	return x, y + 1
}

func (s *state) moves(b bool) []string {
	var result []string

	var x, y int
	if b {
		x, y = s.xr, s.yr
	} else {
		x, y = s.xg, s.yg
	}

	if s.isValid(up(x, y)) {
		result = append(result, "UP")
	}

	if s.isValid(down(x, y)) {
		result = append(result, "DOWN")
	}

	if s.isValid(right(x, y)) {
		result = append(result, "RIGHT")
	}

	if s.isValid(left(x, y)) {
		result = append(result, "LEFT")
	}

	return result
}

func (s *state) clone() state {
	var clone state
	clone.xr = s.xr
	clone.yr = s.yr
	clone.xg = s.xg
	clone.yg = s.yg
	clone.m = make([][]string, 15)
	for i := range clone.m {
		clone.m[i] = make([]string, 15)
		copy(clone.m[i], s.m[i])
	}

	return clone
}

func (s *state) update() {
	s.m[s.xr][s.yr] = "r"
	s.m[s.xg][s.yg] = "g"
}

func (s *state) result(move string, maximize bool) state {
	result := s.clone()

	if maximize {
		switch move {
		case "UP":
			result.xr--
		case "DOWN":
			result.xr++
		case "LEFT":
			result.yr--
		case "RIGHT":
			result.yr++
		}
	} else {
		switch move {
		case "UP":
			result.xg--
		case "DOWN":
			result.xg++
		case "RIGHT":
			result.yg++
		case "LEFT":
			result.yg--
		}
	}

	result.update()
	result.print()
	return result
}

func (s *state) eval(b bool) int {
	if s.xr == s.xg && s.yr == s.yg {
		if b {
			return 50
		}
		return -50
	}

	if s.xr == s.xg {
		d := s.yr - s.yg
		if d < -1 || d > 1 {
			if b {
				return 100 //s.space(s.xr, s.yr)
			}
			return -100
		}
	}

	if s.yr == s.yg {
		d := s.xr - s.xg
		if d < -1 || d > 1 {
			if b {
				return 100 //space(s.xr, s.yr)
			}
			return -100
		}
	}

	return -100
}

// TODO : check booleans
func minimax(s state, depth int, maximize bool) (int, string) {
	moves := s.moves(maximize)
	if depth == 0 || len(moves) == 0 {
		v := s.eval(maximize)
		fmt.Printf("Score : %d\n", v)
		return v, ""
	}

	var score int
	var move string
	if maximize {
		score, move = -1000, "NULL"
		for _, v := range moves {
			fmt.Println(v)
			a, _ := minimax(s.result(v, maximize), depth-1, maximize)
			if a > score {
				score, move = a, v
			}
		}
	} else {
		score, move = 1000, "NULL"
		for _, v := range moves {
			fmt.Println(v)
			a, _ := minimax(s.result(v, !maximize), depth-1, !maximize)
			if a < score {
				score, move = a, v
			}
		}
	}

	fmt.Printf("%v chooses %s with a score of %d\n", maximize, move, score)
	return score, move
}

// Tron bla bla bla
func Tron(inputFile string) {

	input, _ := os.Open(inputFile)
	defer input.Close()

	var p string
	fmt.Fscanf(input, "%s\n", &p)

	var xr, yr, xg, yg int
	fmt.Fscanf(input, "%d %d %d %d\n", &xr, &yr, &xg, &yg)

	m := make([][]string, 15)
	for i := 0; i < 15; i++ {
		var s string
		fmt.Fscanf(input, "%s\n", &s)
		m[i] = strings.Split(s, "")
	}

	//fmt.Println(m)

	if p == "l" {
		xr, yr, xg, yg = xg, yg, xr, yr
	}

	var s state
	s.xr, s.yr, s.xg, s.yg = xr, yr, xg, yg
	s.m = m

	s.print()

	depth := 5
	_, result := minimax(s, depth, true)
	fmt.Println(result)
}
