package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic("Not able to open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var lines []string

	var left []int
	var right []int
	var rightMap = make(map[int]int)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
		tmp := strings.Split(line, "   ")

		l, _ := strconv.Atoi(tmp[0])
		r, _ := strconv.Atoi(tmp[1])
		left = append(left, l)
		right = append(right, r)
	}

	sort.Ints(left)
	sort.Ints(right)

	// fmt.Println(left, right)

	var partOneResult int
	var partTwoResult int

	for i := 0; i < len(left); i++ {

		if left[i] <= right[i] {
			partOneResult += right[i] - left[i]
		} else {
			partOneResult += left[i] - right[i]
		}

		// constructiong a frequncy distribution of right slice

		val, ok := rightMap[right[i]]

		if !ok {
			rightMap[right[i]] = 1
		} else {
			rightMap[right[i]] = val + 1
		}

	}

	fmt.Println(partOneResult)
	// fmt.Println(rightMap)

	for _, rv := range left {
		partTwoResult += rv * rightMap[rv]
	}

	fmt.Println(partTwoResult)

}
