package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	grid := make([][]byte, 0)
	moves := [8][2]int{
		{0, -1},
		{1, 0},
		{0, 1},
		{-1, 0},
		{1, 1},
		{-1, 1},
		{1, -1},
		{-1, -1},
	}
	for _, line := range strings.Split(string(buf), "\n") {
		grid = append(grid, []byte(line))
	}
	// visited grid is used to mark places for numbers that are adjacent to some symbole
	// later, we iterate over the grid but only see the visited cells, to gather/sum all the numbers
	visited := make([][]bool, len(grid))
	sum := 0
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	n, m := len(grid), len(grid[0])
	for i, line := range grid {
		for j, ch := range line {
			ch := byte(ch)
			if ch == '.' || (ch >= '0' && ch <= '9') {
				continue
			}
			// dfs like
			for _, move := range moves {
				_i, _j := i+move[0], j+move[1]
				if _i >= n || _i < 0 || _j >= m || _j < 0 {
					continue
				}
				_ch := grid[_i][_j]
				if _ch < '0' || _ch > '9' {
					continue
				}
				visited[_i][_j] = true
				// consume all digits to the right
				for k := 1; k+_j < m; k++ {
					_ch := grid[_i][_j+k]
					if _ch < '0' || _ch > '9' || visited[_i][_j+k] {
						break
					}
					visited[_i][_j+k] = true
				}
				// consume all digits to the left
				for k := -1; _j+k >= 0; k-- {
					_ch := grid[_i][_j+k]
					if _ch < '0' || _ch > '9' || visited[_i][_j+k] {
						break
					}
					visited[_i][_j+k] = true
				}
			}
		}
	}
	for i, line := range grid {
		acc := ""
		for j, ch := range line {
			if visited[i][j] {
				acc += string(ch)
			} else if len(acc) != 0 {
				// fmt.Println(acc)
				num, _ := strconv.Atoi(acc)
				sum += num
				acc = ""
			}
		}
		if len(acc) != 0 {
			// fmt.Println(acc)
			num, _ := strconv.Atoi(acc)
			sum += num
		}
	}
	fmt.Println(sum)
}
