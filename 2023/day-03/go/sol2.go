package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type num_ptr struct {
	line, st, en int
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
	n, m, sum := len(grid), len(grid[0]), 0
	for i, line := range grid {
		for j, ch := range line {
			ch := byte(ch)
			if ch != '*' {
				continue
			}
			set := make(map[num_ptr]bool)
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
				// l, st, en is used as a view_ptr to the part of the grid that represents a number
				l, st, en, k := _i, _j, _j, 0
				// consume all digits to the right and record the index as `st`
				for k = 1; k+_j < m; k++ {
					_ch := grid[_i][_j+k]
					if _ch < '0' || _ch > '9' {
						break
					}
				}
				en, k = _j+k-1, _j
				// consume all digits to the left and record the index as `en`
				for k = -1; _j+k >= 0; k-- {
					_ch := grid[_i][_j+k]
					if _ch < '0' || _ch > '9' {
						break
					}
				}
				st = _j + k + 1
				// add the number pointer to the set of adjacent numbers
				set[num_ptr{l, st, en}] = true
			}
			product := 1
			if len(set) == 2 {
				for ptr := range set {
					num, err := strconv.Atoi(string(grid[ptr.line][ptr.st : ptr.en+1]))
					if err != nil {
						panic(err)
					}
					product *= num
				}
				sum += product
			}
		}
	}

	fmt.Println(sum)
}
