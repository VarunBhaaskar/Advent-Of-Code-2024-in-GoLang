package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	path := "input.txt"
	input := parse(path)
	fmt.Println(len(input))

	var result int
	var result2 int

	for k, v := range input {
		if isItPossible(v, k[1]) {
			result += k[1]
		}
	}

	for k, v := range input {
		if isItPossible2(v, k[1]) {
			result2 += k[1]
		}
	}

	fmt.Println("Part 1 result: ", result)
	fmt.Println("Part 2 result: ", result2)

}

func parse(path string) map[[2]int][]int {
	input := make(map[[2]int][]int)
	f, err := os.Open(path)

	if err != nil {
		panic("cant open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	index := 0 // had to introduce index because map[int][]int had duplicate keys
	for scanner.Scan() {
		line := scanner.Text()

		split_1 := strings.Split(line, ":")
		split_2 := strings.Split(strings.TrimSpace(split_1[1]), " ")
		// fmt.Println(split_1[0], split_2)

		k, _ := strconv.Atoi(split_1[0])

		tmp := make([]int, len(split_2))

		for i, v := range split_2 {
			v_i, _ := strconv.Atoi(v)
			tmp[i] = v_i
		}

		input[[2]int{index, k}] = tmp
		index++
	}

	return input

}

func isItPossible(input []int, result int) bool {
	// fmt.Println(input, result)
	if len(input) == 1 {
		if input[0] == result {
			// fmt.Println("Match for ", result)
			return true
		} else {
			return false
		}
	} else if len(input) == 2 {
		a1 := isItPossible([]int{input[0] + input[1]}, result)

		if a1 {
			return a1
		} else {
			return isItPossible([]int{input[0] * input[1]}, result)
		}

	} else {

		a2 := isItPossible(append([]int{input[0] + input[1]}, input[2:]...), result)
		if a2 {
			return a2
		} else {
			return isItPossible(append([]int{input[0] * input[1]}, input[2:]...), result)
		}

	}
}

func isItPossible2(input []int, result int) bool {
	// fmt.Println(input, result)
	if len(input) == 1 {
		if input[0] == result {
			// fmt.Println("Match for ", result)
			return true
		} else {
			return false
		}
	} else if len(input) == 2 {
		a1 := isItPossible2([]int{input[0] + input[1]}, result)

		if a1 {
			return a1
		} else {
			a2 := isItPossible2([]int{input[0] * input[1]}, result)
			if a2 {
				return a2
			} else {
				i1, _ := strconv.Atoi(fmt.Sprintf("%s%s", strconv.Itoa(input[0]), strconv.Itoa(input[1])))
				// fmt.Println("len 2", i1)
				return isItPossible2([]int{i1}, result)
			}
		}

	} else {

		if isItPossible2(append([]int{input[0] + input[1]}, input[2:]...), result) {
			return true
		}

		if isItPossible2(append([]int{input[0] * input[1]}, input[2:]...), result) {
			return true
		}

		i1, _ := strconv.Atoi(fmt.Sprintf("%s%s", strconv.Itoa(input[0]), strconv.Itoa(input[1])))
		// fmt.Println(i1)
		return isItPossible2(append([]int{i1}, input[2:]...), result)

	}
}
