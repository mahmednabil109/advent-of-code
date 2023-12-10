package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func endCheck(values []int) bool {
	start := values[0]
	for _, v := range values[1:] {
		if start != v {
			return false
		}
	}
	return true
}

func predictPrevValue(history []int) int {
	pascal := [][]int{
		history,
	}
	for i := 0; !endCheck(pascal[i]); i++ {
		next := make([]int, 0)
		for j, v := range pascal[i][1:] {
			next = append(next, v-pascal[i][j])
		}
		pascal = append(pascal, next)
	}
	n := len(pascal) - 1
	ans := pascal[n][0]
	for i := n - 1; i >= 0; i-- {
		ans = pascal[i][0] - ans
	}
	return ans
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*22)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	ans := 0
	for _, line := range strings.Split(string(buf), "\n") {
		data := strings.Split(line, " ")
		history := make([]int, 0)
		for _, d := range data {
			num, err := strconv.Atoi(d)
			if err != nil {
				panic(err)
			}
			history = append(history, num)
		}
		ans += predictPrevValue(history)
	}
	fmt.Println(ans)
}
