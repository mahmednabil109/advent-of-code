package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Trie struct {
	next  [26]*Trie
	value int
	end   bool
}

func NewTrie(keys []string) *Trie {
	t := Trie{}
	for i := 0; i < len(keys); i++ {
		t.insert(keys[i], i)
	}
	return &t
}

func (t *Trie) insert(key string, value int) {
	curr := t
	for pos, ch := range key {
		idx := int(byte(ch) - 'a')
		if curr.next[idx] == nil {
			curr.next[idx] = &Trie{}
		}
		curr = curr.next[idx]
		if pos == len(key)-1 {
			curr.value = value
			curr.end = true
		}
	}
}

func (t *Trie) Next(ch byte) (*Trie, bool) {
	idx := int(ch - 'a')
	if t.next[idx] == nil {
		return nil, false
	}
	return t.next[idx], true
}

func (t *Trie) Search(key string) (int, bool) {
	curr := t
	for _, ch := range key {
		idx := int(byte(ch) - 'a')
		if idx < 0 || idx > 26 {
			break
		}
		if curr.next[idx] == nil {
			return 0, false
		}
		curr = curr.next[idx]
		if curr.end {
			return curr.value, true
		}
	}
	return curr.value, curr.end
}

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
	trie := NewTrie([]string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	})

	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		var fd, ld, d int
		first, ok := true, false

		for i := 0; i < len(line); i++ {
			ch := line[i]
			if ch >= '0' && ch <= '9' {
				d, ok = int(ch-'0'), true
			} else {
				d, ok = trie.Search(line[i:])
			}
			if ok {
				if first {
					first = false
					fd = d
				}
				ld = d
			}
			ok = false
		}
		// fmt.Printf("[%v]: %v %v\n", line, fd, ld)
		sum += fd*10 + ld
	}

	fmt.Printf("Sum is %v\n", sum)
}
