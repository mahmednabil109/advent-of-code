package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Galaxy struct {
	x, y int64
}

func Max(arr ...int64) int64 {
	max := arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func Abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func eculidianDistance(galaxy1, galaxy2 *Galaxy) int64 {
	return Abs(galaxy1.x-galaxy2.x) + Abs(galaxy1.y-galaxy2.y)
}

func expandColumns(galaxies []*Galaxy, expandRatio int64) {
	sort.Slice(galaxies, func(i, j int) bool {
		return galaxies[i].y < galaxies[j].y
	})
	expand := galaxies[0].y * (expandRatio - 1)
	galaxies[0].y += expand
	for i, galaxy := range galaxies[1:] {
		expand += Max(0, galaxy.y-galaxies[i].y-1+expand) * (expandRatio - 1)
		galaxy.y += expand
	}
}

func expandRows(galaxies []*Galaxy, expandRatio int64) {
	sort.Slice(galaxies, func(i, j int) bool {
		return galaxies[i].x < galaxies[j].x
	})
	expand := galaxies[0].x
	galaxies[0].x += expand * (expandRatio - 1)
	for i, galaxy := range galaxies[1:] {
		expand += Max(0, galaxy.x-galaxies[i].x-1+expand) * (expandRatio - 1)
		galaxy.x += expand
	}
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*20)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]
	galaxies := make([]*Galaxy, 0)
	for i, line := range strings.Split(string(buf), "\n") {
		for j, c := range line {
			if c == '#' {
				galaxies = append(galaxies, &Galaxy{x: int64(i), y: int64(j)})
			}
		}
	}

	expandColumns(galaxies, 1_000_000)
	expandRows(galaxies, 1_000_000)

	ans := int64(0)
	for i, galaxy := range galaxies {
		for _, galaxy2 := range galaxies[i+1:] {
			ans += eculidianDistance(galaxy, galaxy2)
		}
	}
	fmt.Println(ans)

}
