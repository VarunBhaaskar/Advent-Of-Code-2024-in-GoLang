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

	fmt.Println("Part-1: ", part1(path))
	fmt.Println("Part-2: ", part2(path))
}

func parseInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		panic("Cannot open file")
	}

	return string(data)
}

func mul(a string) int {

	r, _ := regexp.Compile(`(\d{1,},\d{1,})`)
	nums := strings.Split(r.FindString(a), ",")

	a1, _ := strconv.Atoi(nums[0])
	a2, _ := strconv.Atoi(nums[1])
	return a1 * a2
}

func part1(path string) int {
	var result int
	r := regexp.MustCompile(`(?m)mul\(\d{1,},\d{1,}\)`)
	fmt.Println(r)

	input := parseInput(path)
	// fmt.Println(input)
	matches := r.FindAllString(input, -1)

	fmt.Println(matches, len(matches))

	for _, v := range matches {
		result += mul(v)
	}

	return result

}

func part2(path string) int {
	var result int
	var mulEnabled int = 1
	r := regexp.MustCompile(`((?m)mul\(\d{1,},\d{1,}\))|(do\(\)|don't\(\))`)
	fmt.Println(r)

	input := parseInput(path)
	// fmt.Println(input)
	matches := r.FindAllString(input, -1)

	fmt.Println(matches, len(matches))
	for _, v := range matches {
		if v == "do()" {
			mulEnabled = 1
			continue
		} else if v == "don't()" {
			mulEnabled = 0
			continue
		} else {
			result += mulEnabled * mul(v)
		}
	}
	return result
}
