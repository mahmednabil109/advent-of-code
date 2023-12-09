package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcSteps(adjList map[string][2]string, inst string) int {
	steps, curr := 0, "AAA"
	for i := 0; ; i = (i + 1) % len(inst) {
		if curr == "ZZZ" {
			break
		}
		move := 0
		if inst[i] == 'R' {
			move = 1
		}
		curr = adjList[curr][move]
		steps += 1
	}

	return steps
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*14)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	lines := strings.Split(string(buf), "\n")
	instructions := lines[0]

	adjList := make(map[string][2]string)
	for _, line := range lines[2:] {
		parts := strings.Split(line, "=")
		node, list := strings.Trim(parts[0], " "), strings.Split(parts[1], ",")
		adjList[node] = [2]string{
			list[0][2:], list[1][1 : len(list[1])-1],
		}
	}
	fmt.Println(calcSteps(adjList, instructions))
}
