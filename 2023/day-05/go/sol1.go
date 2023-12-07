package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mapConv struct {
	Ranges []struct {
		destSt, sourSt, l int
	}
}

func (m mapConv) Convert(num int) int {
	for _, r := range m.Ranges {
		if num >= r.sourSt && num < (r.sourSt+r.l) {
			return num + (r.destSt - r.sourSt)
		}
	}
	return num
}

func filterLines(data []string) (result []string) {
	result = make([]string, 0, len(data))
	for _, l := range data {
		if len(l) == 0 || l[0] < '0' || l[0] > '9' {
			continue
		}
		result = append(result, l)
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
	buf := make([]byte, 1024*5)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	parts := strings.Split(string(buf), ":")
	// parse seeds
	seeds := make([]int, 0)
	for _, s := range strings.Split(
		strings.Trim(strings.Split(parts[1], "\n")[0], " \n"),
		" ",
	) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, num)
	}

	// parse maps
	converts := make([]mapConv, 0)
	for _, m := range parts[2:] {
		ranges := filterLines(strings.Split(m, "\n"))
		mc := mapConv{}
		for _, r := range ranges {
			data := strings.Split(strings.Trim(r, " \n"), " ")
			intdata := make([]int, 0, len(data))
			for _, d := range data {
				num, err := strconv.Atoi(d)
				if err != nil {
					panic(err)
				}
				intdata = append(intdata, num)
			}
			mc.Ranges = append(mc.Ranges, struct {
				destSt int
				sourSt int
				l      int
			}{intdata[0], intdata[1], intdata[2]})
		}
		converts = append(converts, mc)
	}
	ans := -1
	for _, seed := range seeds {
		res := seed
		for _, conv := range converts {
			res = conv.Convert(res)
		}
		if ans == -1 || ans > res {
			ans = res
		}
	}
	fmt.Println(ans)
}
