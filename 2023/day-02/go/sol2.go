package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Max(arr ...int) (res int) {
	res = arr[0]
	for _, el := range arr[1:] {
		if el > res {
			res = el
		}
	}
	return
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 11*1024)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	sum := 0
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		colon_end := strings.Index(line, ":")
		sets := strings.Split(line[colon_end+1:], ";")
		hmap := make(map[string]int)
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				data := strings.Split(strings.Trim(cube, " "), " ")
				count, _ := strconv.Atoi(data[0])
				hmap[data[1]] = Max(hmap[data[1]], count)
			}
		}
		if len(hmap) > 3 {
			panic("there should be at most 3")
		}
		power := 1
		for _, value := range hmap {
			power *= value
		}

		sum += power
		// fmt.Printf("Game: %v \n", power)
	}

	fmt.Println(sum)
}
