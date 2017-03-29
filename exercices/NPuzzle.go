package exercices

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func newState(size int) [][]int {
	result := make([][]int, size)
	for i := 0; i < size; i++ {
		result[i] = make([]int, size)
	}
	return result
}

type coords struct {
	x, y int
}

type node struct {
	parent   *node
	state    [][]int
	action   string
	pathCost int
}

// ByPathCost sort frontier by path-cost
type ByPathCost []node

func (bpc ByPathCost) Len() int {
	return len(bpc)
}

func (bpc ByPathCost) Less(i, j int) bool {
	return bpc[i].pathCost < bpc[j].pathCost
}

func (bpc ByPathCost) Swap(i, j int) {
	bpc[i], bpc[j] = bpc[j], bpc[i]
}

// test if 2 states are equal
func equal(s1, s2 [][]int) bool {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1[i]); j++ {
			if s1[i][j] != s2[i][j] {
				return false
			}
		}
	}

	return true
}

func actions(s *[][]int) []string {
	var zero coords
	for i := 0; i < len(*s); i++ {
		for j := 0; j < len((*s)[i]); j++ {
			if (*s)[i][j] == 0 {
				zero.x, zero.y = i, j
			}
		}
	}

	var result []string
	if zero.x-1 >= 0 {
		result = append(result, "UP")
	}

	if zero.y-1 >= 0 {
		result = append(result, "LEFT")
	}

	if zero.y+1 < len((*s)[0]) {
		result = append(result, "RIGHT")
	}

	if zero.x+1 < len(*s) {
		result = append(result, "DOWN")
	}

	return result
}

func inExplored(a *[][][]int, s *[][]int) bool {
	i := 0
	for i < len(*a) {
		if equal((*a)[i], *s) {
			return true
		}
		i++
	}

	return false
}

func inFrontier(a *[]node, s *[][]int) (bool, int) {
	i := 0
	for i < len(*a) {
		if equal((*a)[i].state, *s) {
			return true, i
		}
		i++
	}

	return false, 0
}

func childNode(n node, m string) node {
	var result node
	result.parent = &n
	result.state = apply(&(n.state), m)
	result.action = m

	result.pathCost = n.pathCost + 1 + h(&(result.state))

	return result
}

func h(s *[][]int) int {
	result := 0
	for i := range *s {
		for j := range (*s)[i] {
			switch (*s)[i][j] {
			case 1:
				result += int(math.Abs(float64(0)-float64(i)) + math.Abs(float64(1)-float64(j)))
			case 2:
				result += int(math.Abs(float64(0)-float64(i)) + math.Abs(float64(2)-float64(j)))
			case 3:
				result += int(math.Abs(float64(1)-float64(i)) + math.Abs(float64(0)-float64(j)))
			case 4:
				result += int(math.Abs(float64(1)-float64(i)) + math.Abs(float64(1)-float64(j)))
			case 5:
				result += int(math.Abs(float64(1)-float64(i)) + math.Abs(float64(2)-float64(j)))
			case 6:
				result += int(math.Abs(float64(2)-float64(i)) + math.Abs(float64(0)-float64(j)))
			case 7:
				result += int(math.Abs(float64(2)-float64(i)) + math.Abs(float64(1)-float64(j)))
			case 8:
				result += int(math.Abs(float64(2)-float64(i)) + math.Abs(float64(2)-float64(j)))
			}
		}
	}
	return result
}

func apply(s *[][]int, m string) [][]int {
	result := make([][]int, len(*s))
	for i := range result {
		result[i] = make([]int, len((*s)[i]))
	}

	var zero coords
	for i := range *s {
		for j := range (*s)[i] {
			result[i][j] = (*s)[i][j]
			if result[i][j] == 0 {
				zero.x, zero.y = i, j
			}
		}
	}

	switch m {
	case "UP":
		{
			result[zero.x][zero.y], result[zero.x-1][zero.y] = result[zero.x-1][zero.y], result[zero.x][zero.y]
		}
	case "LEFT":
		{
			result[zero.x][zero.y], result[zero.x][zero.y-1] = result[zero.x][zero.y-1], result[zero.x][zero.y]
		}
	case "RIGHT":
		{
			result[zero.x][zero.y], result[zero.x][zero.y+1] = result[zero.x][zero.y+1], result[zero.x][zero.y]
		}
	case "DOWN":
		{
			result[zero.x][zero.y], result[zero.x+1][zero.y] = result[zero.x+1][zero.y], result[zero.x][zero.y]
		}
	}

	return result
}

func solution(n node) {
	c := n
	var path []string
	for c.parent != nil {
		path = append(path, c.action)
		c = *(c.parent)
	}

	fmt.Println(len(path))
	i := len(path) - 1
	for i >= 0 {
		fmt.Println(path[i])
		i--
	}
}

// NPuzzle bla bla bla
func NPuzzle(inputFile string) {
	input, _ := os.Open(inputFile)
	defer input.Close()

	var size int
	fmt.Fscanf(input, "%d\n", &size)

	var zero coords
	initialState := newState(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Fscanf(input, "%d\n", &initialState[i][j])
			if initialState[i][j] == 0 {
				zero.x, zero.y = i, j
			}
		}
	}

	goal := newState(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			goal[i][j] = (size*i + j)
		}
	}

	var frontier []node
	var explored [][][]int

	var start node
	start.state = initialState
	start.parent = nil
	start.action = ""
	start.pathCost = 0

	frontier = append(frontier, start)
	var n node
	for len(frontier) > 0 {
		n, frontier = frontier[0], frontier[1:]
		if equal(n.state, goal) {
			solution(n)
			break
		} else {
			explored = append(explored, n.state)
			moves := actions(&(n.state))
			for _, m := range moves {
				child := childNode(n, m)

				if !inExplored(&explored, &(child.state)) {
					b, k := inFrontier(&frontier, &(child.state))
					if !b {
						frontier = append(frontier, child)
					} else {
						if child.pathCost < frontier[k].pathCost {
							fmt.Println("Changing", child, frontier[k])
							frontier[k] = child
						}
					}
				}
			}

			sort.Sort(ByPathCost(frontier))
		}
	}
}
