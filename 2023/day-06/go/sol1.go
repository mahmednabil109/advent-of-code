package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func splitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	i, seen, last := 0, false, true
	for j, b := range data {
		i = j
		if b == ' ' && !seen {
			continue
		} else if b != ' ' {
			seen = true
		} else {
			last = false
			break
		}
	}
	if i >= 0 && !last {
		// fmt.Println(i, string(data[:i]))
		return i + 1, data[:i], nil
	}
	if atEOF {
		return len(data), data, nil
	}
	return 0, nil, nil
}

func solveQuadratic(T, D int) int {
	criticalTime := (float64(T) - math.Sqrt(float64(T*T-4*D))) / 2.0
	return int(math.Floor(criticalTime + 1))
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	lines := strings.Split(strings.Trim(string(buf), " \n"), "\n")
	ans := 1
	times := make([]int, 0)
	scanner := bufio.NewScanner(
		bytes.NewBufferString(lines[0][5:]),
	)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		num, err := strconv.Atoi(strings.Trim(scanner.Text(), " \n"))
		if err != nil {
			panic(err)
		}
		times = append(times, num)
	}
	distances := make([]int, 0)
	scanner = bufio.NewScanner(
		bytes.NewBufferString(lines[1][9:]),
	)
	scanner.Split(splitFunc)
	for scanner.Scan() {
		num, err := strconv.Atoi(strings.Trim(scanner.Text(), " \n"))
		if err != nil {
			panic(err)
		}
		distances = append(distances, num)
	}
	for i, time := range times {
		winningTime := solveQuadratic(time, distances[i])
		ans *= (time - 2*winningTime + 1)
	}

	fmt.Println(ans)
}
