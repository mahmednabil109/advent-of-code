package main

import (
	"bufio"
	"bytes"
	"fmt"
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

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*25)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	sum := 0
	lines := strings.Split(string(buf), "\n")
	cards := make([]int, len(lines))
	for i := range cards {
		cards[i] = 1
	}

	for i, line := range lines {
		colon_idx := strings.Index(line, ":")
		line = line[colon_idx+1:]
		sets := strings.Split(line, "|")

		scanner := bufio.NewScanner(
			bytes.NewBufferString(sets[0]),
		)
		scanner.Split(splitFunc)
		// scan winning numbers
		winner := make(map[int]bool)
		// fmt.Println(sets[0])
		for scanner.Scan() {
			num, err := strconv.Atoi(strings.Trim(scanner.Text(), " "))
			if err != nil {
				panic(err)
			}
			winner[num] = true
		}

		scanner = bufio.NewScanner(
			bytes.NewBufferString(sets[1]),
		)
		scanner.Split(splitFunc)
		// scan numbers
		// fmt.Println(sets[1])
		count := 0
		for scanner.Scan() {
			num, err := strconv.Atoi(strings.Trim(scanner.Text(), " "))
			if err != nil {
				panic(err)
			}
			if _, ok := winner[num]; ok {
				count += 1
			}
		}
		for j := 1; j <= count; j++ {
			cards[i+j] += cards[i]
		}
	}
	for _, count := range cards {
		sum += count
	}
	fmt.Println(sum)
}
