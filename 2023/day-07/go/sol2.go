package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand string
	bid  int
}

func getType(hand string) int {
	set, first := make(map[byte]int), hand[0]
	for _, r := range hand {
		set[byte(r)] += 1
	}
	// adapt J
	if J_cnt, ok := set['J']; ok {
		max_k, max_cnt := byte('$'), 0
		for k, cnt := range set {
			if cnt > max_cnt && k != 'J' {
				max_k = k
				max_cnt = cnt
			}
		}
		set[max_k] += J_cnt
		first = max_k
		delete(set, 'J')
	}
	switch l := len(set); l {
	case 1:
		return 7
	case 2:
		if cnt := set[first]; cnt == 1 || cnt == 4 {
			return 6
		}
		return 5
	case 3:
		for _, r := range hand {
			if cnt := set[byte(r)]; cnt == 3 {
				return 4
			}
		}
		return 3
	default:
		return 6 - l
	}
}

var powers map[byte]int = map[byte]int{
	'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10, '9': 9, '8': 8, '7': 7, '6': 6, '5': 5, '4': 4, '3': 3, '2': 2,
}

func getPower(c byte) int {
	return powers[c]
}

func main() {
	file, err := os.Open("./in.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buf := make([]byte, 1024*10)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	buf = buf[:n]

	hands := make([]Hand, 0)
	for _, line := range strings.Split(string(buf), "\n") {
		parts := strings.Split(line, " ")
		hand := parts[0]
		bid, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		hands = append(hands, Hand{hand, bid})
	}
	sort.Slice(hands, func(i, j int) bool {
		r1, r2 := getType(hands[i].hand), getType(hands[j].hand)
		// fmt.Println(r1, hands[i], r2, hands[j])
		if r1 != r2 {
			return r1 < r2
		}
		h1, h2 := hands[i].hand, hands[j].hand
		for i := 0; i < 5; i++ {
			c1, c2 := byte(h1[i]), byte(h2[i])
			if c1 != c2 {
				return getPower(c1) < getPower(c2)
			}
		}
		// should not reach here
		panic("no no")
		return false
	})
	ans := int64(0)
	for i, hand := range hands {
		ans += int64(i+1) * int64(hand.bid)
		fmt.Println(getType(hand.hand), hand.hand, hand.bid)
	}
	fmt.Println(ans)
}
