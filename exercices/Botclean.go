package exercices

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type square struct {
	id, x, y int
}

func (c square) distFrom(d square) int {
	return int(math.Abs(float64(c.x)-float64(d.x)) + math.Abs(float64(c.y)-float64(d.y)))
}

type pair struct {
	id, cost int
}

// sort array of pairs by distance
type ByDistance []pair

func (b ByDistance) Len() int {
	return len(b)
}

func (b ByDistance) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByDistance) Less(i, j int) bool {
	return b[i].cost < b[j].cost
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func dfs(to square, dirties *[]square, distances *[][]pair, nb int, visited *[]bool) int {
	fmt.Printf("NB = %d\n", nb)

	if nb == 2 {
		return 0
	}

	a := *dirties
	b := *distances
	c := *visited

	c[to.id] = true

	result := 0
	for _, v := range b[to.id][:min(3, len(b[to.id]))] {
		if !c[v.id] {
			fmt.Println("visiting", a[v.id], "from", a[to.id])
			fmt.Println(b)
			r := b[to.id][v.id].cost
			t := dfs(a[v.id], dirties, distances, nb+1, visited)
			fmt.Printf("r=%d t=%d result=%d\n", r, t, result)
			if result == 0 || r+t < result {
				result = r + t
			}
		}
	}

	c[to.id] = false
	return result
}

func Botclean(bx, by, h, w int, b [][]string) string {
	fmt.Println("botclean", bx, by)

	if b[bx][by] == "d" {
		return "CLEAN"
	}

	// find all dirty squares
	dirties := make([]square, 0, h*w)
	id := 0
	for i, r := range b {
		for j := range r {
			if b[i][j] == "d" {
				dirties = append(dirties, square{id, i, j})
				id++
			}
		}
	}

	fmt.Fprintln(os.Stderr, dirties)

	// compute distances
	distances := make([][]pair, len(dirties))
	fromBot := make([]pair, len(dirties))

	bot := square{0, bx, by}
	for i, v := range dirties {
		fromBot[i] = pair{i, bot.distFrom(v)}

		distances[i] = make([]pair, len(dirties))
		for j, v2 := range dirties {
			distances[i][j] = pair{v2.id, v.distFrom(v2)}
		}
		sort.Sort(ByDistance(distances[i]))
	}

	sort.Sort(ByDistance(fromBot))

	//fmt.Fprintln(os.Stderr, distances)
	//fmt.Fprintln(os.Stderr, fromBot)

	d := 1000
	resx, resy := 0, 0
	for _, v := range fromBot {
		visited := make([]bool, len(fromBot))
		r := v.cost
		t := dfs(dirties[v.id], &dirties, &distances, 1, &visited)
		if r+t < d {
			d = r + t
			resx = dirties[v.id].x
			resy = dirties[v.id].y
		}
	}

	if resx < bx {
		return "UP"
	} else if resx > bx {
		return "DOWN"
	} else if resy < by {
		return "LEFT"
	}

	return "RIGHT"
}
