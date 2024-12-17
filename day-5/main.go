package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	path := "input.txt"
	rules, input := parse(path)
	var result int
	var result2 int
	for _, v := range input {
		res, ok := check(rules, v)

		if ok {
			result += res
		} else {
			result2 += correct(rules, v)
		}
	}

	fmt.Println("part 1 result: ", result)
	fmt.Println("Part 2 result: ", result2)
}

func parse(path string) (map[int][]int, [][]int) {

	rules := make(map[int][]int)
	var input [][]int

	f, err := os.ReadFile(path)
	if err != nil {
		panic("Cannot open file")
	}

	var input_str string = string(f)
	// fmt.Println(input_str)
	input_slice := strings.Split(input_str, "\n")

	regex, _ := regexp.Compile(`^[\t\r\n]+`)

	input_k := 0
	for k, v := range input_slice {

		cutoff := regex.FindStringIndex(v)
		if len(cutoff) > 0 {
			input_k = k
			break
		}

	}
	fmt.Println(input_k)

	for _, v := range input_slice[0:input_k] {

		int_slc := strings.Split(strings.TrimSpace(v), "|")
		// fmt.Println(int_slc[0], int_slc[1])
		v1, _ := strconv.Atoi(int_slc[0])
		v2, _ := strconv.Atoi(int_slc[1])

		_, ok := rules[v1]
		if ok {
			rules[v1] = append(rules[v1], v2)
		} else {
			rules[v1] = []int{v2}
		}

	}

	fmt.Println(rules)

	for _, v := range input_slice[input_k+1:] {
		int_slc := strings.Split(strings.TrimSpace(v), ",")
		// fmt.Println("int slc", int_slc)
		var tmp []int

		for _, i := range int_slc {
			iv, err := strconv.Atoi(i)

			if err != nil {
				fmt.Println(err)
			} else {
				tmp = append(tmp, iv)
			}
		}

		input = append(input, tmp)
	}

	// fmt.Println(input)

	return rules, input
}

func check(rules map[int][]int, arr []int) (int, bool) {
	var forbidden []int
	n := len(arr)

	for i := n - 1; i >= 0; i-- {

		if search(forbidden, arr[i]) {
			fmt.Println("Forbidden", arr)
			return 0, false
		}

		forbiddenElements, ok := rules[arr[i]]
		// fmt.Println("Rules", forbiddenElements, ok, i)
		if ok {
			forbidden = append(forbidden, forbiddenElements...)
		}

		// fmt.Println("New forbidden list for: ", arr[i], forbidden)

	}
	// fmt.Println(arr, arr[(n/2)])
	return arr[(n / 2)], true
}

func search(arr []int, k int) bool {
	for _, v := range arr {
		if v == k {
			return true
		}
	}

	return false
}

func search2(arr []int, k int) (int, bool) {
	for i, v := range arr {
		if v == k {
			return i, true
		}
	}

	return 0, false
}

func correct(rules map[int][]int, arr []int) int {

	var result []int

	for i := 0; i < len(arr); i++ {

		a := len(result)

		if a == 0 {
			result = append(result, arr[i])
		} else {
			position := 0
			for j := 0; j < a; j++ {
				rule, ok := rules[result[j]]

				if ok {
					if search(rule, arr[i]) {
						position = j + 1
						// fmt.Println("New position", arr[i], position, result)
					}
				}
			}

			if position >= a {
				result = append(result, arr[i])
			} else if position == 0 {
				result = append([]int{arr[i]}, result...)
			} else {
				result = append(result[:position+1], result[position:]...)
				result[position] = arr[i]
			}

		}
	}
	fmt.Println("Corrected: ", result)
	fmt.Println(check(rules, result))
	return result[len(result)/2]
}
