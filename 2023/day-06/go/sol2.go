package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solveQuadratic(T, D int64) int64 {
	criticalTime := (float64(T) - math.Sqrt(float64(T*T-4*D))) / 2.0
	return int64(math.Floor(criticalTime + 1))
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
	time, _ := strconv.ParseInt(
		strings.Replace(lines[0][5:], " ", "", -1),
		10,
		64,
	)

	distance, _ := strconv.ParseInt(
		strings.Replace(lines[1][9:], " ", "", -1),
		10,
		64,
	)
	winningTime := solveQuadratic(time, distance)
	ans := (time - 2*winningTime + 1)

	fmt.Println(ans)
}
