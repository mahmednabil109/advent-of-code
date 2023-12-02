package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 21*1024)
	reader := bufio.NewReader(file)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	var sum int
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		var fd, ld int
		first := true
		for _, chr := range line {
			chr := byte(chr)
			if chr >= '0' && chr <= '9' {
				if first {
					fd = int(chr - '0')
					first = false
				}
				ld = int(chr - '0')
			}
		}
		sum += fd*10 + ld
	}

	fmt.Printf("Sum is %v\n", sum)
}
