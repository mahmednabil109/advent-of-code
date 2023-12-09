package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func calcSteps(start string, adjList map[string][2]string, inst string) int {
	steps, curr := 0, start
	for i := 0; ; i = (i + 1) % len(inst) {
		if curr[2] == 'Z' {
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

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
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
	starts := make([]string, 0)
	for _, line := range lines[2:] {
		parts := strings.Split(line, "=")
		node, list := strings.Trim(parts[0], " "), strings.Split(parts[1], ",")
		adjList[node] = [2]string{
			list[0][2:], list[1][1 : len(list[1])-1],
		}
		if node[2] == 'A' {
			starts = append(starts, node)
		}
	}
	steps := make([]int, 0)
	// get the individaul steps,
	// and to reach xxZ with all of them the answer would be the LCM of all steps;
	// as the steps are cyclic; i.e. if we can reach Z from A in N steps, we can do in M * N  steps.
	for _, start := range starts {
		steps = append(steps, calcSteps(start, adjList, instructions))
	}
	// calculate the LCM
	ans := steps[0]
	for _, step := range steps[1:] {
		ans = (step * ans) / gcd(step, ans)
	}
	fmt.Println(ans)
}
