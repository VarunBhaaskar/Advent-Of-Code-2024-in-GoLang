package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	input := []int{0, 7, 198844, 5687836, 58, 2478, 25475, 894}
	// sampleInput := []int{125, 17}
	resultInt := part1(input)
	fmt.Println("Part 1 result: ", resultInt)
	fmt.Println("part 2 result: ", part2(input))
}

func part1(input []int) int {
	var resultInt int
	for _, v := range input {
		resultInt += transform(25, v)
	}

	return resultInt
}

func part2(input []int) int {
	var resultInt int

	var inputMap = make(map[int]int)
	for _, v := range input {
		inputMap[v] = 1
	}
	for i := 0; i < 75; i++ {
		// fmt.Println(inputMap)
		inputMap = transform2(inputMap)
	}

	for _, v := range inputMap {
		resultInt += v
	}
	return resultInt
}

func transform2(inputMap map[int]int) map[int]int {
	var newResult = make(map[int]int)
	for k, v := range inputMap {

		t := nextState(k)

		if t[0] != -1 {
			newResult[t[0]] += v
		}

		if t[1] != -1 {
			newResult[t[1]] += v
		}

		// fmt.Println(result)
	}

	return newResult
}

func transform(n int, val int) int {
	var result []int
	result = append(result, val)

	for i := 0; i < n; i++ {

		var tmp [][2]int
		for _, v := range result {
			t := nextState(v)
			tmp = append(tmp, t)
		}

		result = unpack(tmp)

		// fmt.Println(result)

	}
	return len(result)
}

func nextState(v int) [2]int {

	var res [2]int

	if v == 0 {
		res[0] = 1
		res[1] = -1
	} else if a := len(strconv.Itoa(v)); a%2 == 0 {
		a = a / 2

		d := int(math.Pow(10, float64(a)))

		n1 := v / d
		n2 := v - (n1 * d)
		res[0] = n1
		res[1] = n2
	} else {
		res[0] = v * 2024
		res[1] = -1
	}
	return res
}

func unpack(v [][2]int) []int {
	var result []int

	for _, x := range v {
		if x[0] != -1 {
			result = append(result, x[0])
		}

		if x[1] != -1 {
			result = append(result, x[1])
		}
	}

	return result
}
