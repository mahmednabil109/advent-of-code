package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Min(arr ...int) (res int) {
	res = arr[0]
	for _, el := range arr {
		if el < res {
			res = el
		}
	}
	return
}

type SeedRange struct {
	st, l int
}

type MapRange struct {
	destSt, sourSt, l int
}

type MapConv struct {
	Ranges []MapRange
}

func NewMapConv(ranges []MapRange) MapConv {
	sort.Slice(ranges, func(a, b int) bool {
		return ranges[a].sourSt < ranges[b].sourSt
	})
	return MapConv{
		// mistake; don't do that
		Ranges: ranges,
	}
}

func (m MapConv) Convert(num int) int {
	for _, r := range m.Ranges {
		if num >= r.sourSt && num < (r.sourSt+r.l) {
			return num + (r.destSt - r.sourSt)
		}
	}
	return num
}

func (m MapConv) ConvertRange(input SeedRange) []SeedRange {
	// fmt.Println("\n\n\n", m.Ranges, input)
	result := make([]SeedRange, 0)
	for _, r := range m.Ranges {
		if input.l == 0 {
			break
		}
		switch {
		case input.st < r.sourSt:
			l := Min(input.l, r.sourSt-input.st)
			result = append(result, SeedRange{
				st: input.st,
				l:  l,
			})
			input.st = r.sourSt
			input.l -= l
		case input.st == r.sourSt:
			l := Min(input.l, r.l)
			result = append(result, SeedRange{
				st: r.destSt,
				l:  l,
			})
			input.st += l
			input.l -= l
		default:
			if input.st >= r.sourSt+r.l {
				continue
			}
			l := Min(input.l, r.l-(input.st-r.sourSt))
			result = append(result, SeedRange{
				st: r.destSt + (input.st - r.sourSt),
				l:  l,
			})
			input.st += l
			input.l -= l
		}
		// fmt.Println(input.st, input.l)
	}
	if input.l != 0 {
		result = append(result, SeedRange{
			st: input.st,
			l:  input.l,
		})
	}
	return result
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
	converts := make([]MapConv, 0)
	for _, m := range parts[2:] {
		ranges := filterLines(strings.Split(m, "\n"))
		mr := make([]MapRange, 0)
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
			mr = append(mr, MapRange{
				destSt: intdata[0],
				sourSt: intdata[1],
				l:      intdata[2],
			})
		}
		mc := NewMapConv(mr)
		converts = append(converts, mc)
	}
	// fmt.Printf("%+v \n", converts[0])
	ans := -1
	for i := 0; i < len(seeds)-1; i += 2 {
		st, l := seeds[i], seeds[i+1]
		res := []SeedRange{{st, l}}
		for _, conv := range converts {
			tmp := make([]SeedRange, 0, len(res))
			for _, sr := range res {
				tmp = append(tmp, conv.ConvertRange(sr)...)
			}
			res = tmp
			// fmt.Println("res", res)
		}
		for _, r := range res {
			if ans == -1 {
				ans = r.st
			}
			ans = Min(ans, r.st)
		}
	}
	fmt.Println(ans)
}
