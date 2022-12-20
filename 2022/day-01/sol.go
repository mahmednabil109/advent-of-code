package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func Max(nums ...int) int{
	res := nums[0]
	for _, num := range nums{
		if res < num {
			res = num
		}
	}
	return res
}

func main(){
	file, err := os.Open("./input1")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	var ans1 int //, ans2 int32
	// var cals_per_elf []int
	reader := bufio.NewReader(file)
	buf := make([]byte, 3024)
	n, err := reader.Read(buf)
	if err != nil {
		panic(err)
	}
	fmt.Printf("bytes read: %v \n", n)
	
	fmt.Println(string(buf))
	for _, elfs := range strings.Split(string(buf), "\n\n") {
		cal_sum := 0
		for _, cal := range strings.Split(elfs, "\n") {
			if cal == "" {
				continue
			}
			cal_int, err := strconv.Atoi(cal)
			if err != nil {
				panic(err)
			}
			cal_sum += cal_int
		}
		// cals_per_elf = append(cals_per_elf, cal_sum)
		ans1 = Max(cal_sum, ans1)
	}

	fmt.Printf("%v \n", ans1)
}
