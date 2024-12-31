package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Equation struct {
	a1, a2, b1, b2, c1, c2 int
}

func main() {
	path := "input.txt"
	var result1 int
	var result2 int
	input := parse(path)

	for _, eq := range input {
		result1 += solveEquation(eq)
	}

	for _, eq := range input {
		eq.c1 = eq.c1 + 10000000000000
		eq.c2 = eq.c2 + 10000000000000
		result2 += solveEquation(eq)
	}

	fmt.Println("Part 1 result : ", result1)
	fmt.Println("Part 2 result : ", result2)
}

func parse(path string) []Equation {
	f, err := os.ReadFile(path)

	if err != nil {
		panic("Cant open file")
	}

	var input []Equation

	f_string := string(f)

	strNormalized := regexp.
		MustCompile("\r\n").
		ReplaceAllString(f_string, "\n")

	f_slice := regexp.
		MustCompile(`\n\s*\n`).
		Split(strNormalized, -1)

	var re = regexp.MustCompile(`(?m)Button A: X\+(\d{1,}), Y\+(\d{1,})\nButton B: X\+(\d{1,}), Y\+(\d{1,})\nPrize: X=+(\d{1,}), Y=(\d{1,})`)

	for _, eq := range f_slice {
		// fmt.Println(eq)
		cooeffs_str := re.FindStringSubmatch(eq)
		var cooeffs []int

		for _, v := range cooeffs_str[1:] {
			vi, _ := strconv.Atoi(v)
			cooeffs = append(cooeffs, vi)
		}

		input = append(input, Equation{cooeffs[0], cooeffs[1], cooeffs[2], cooeffs[3], cooeffs[4], cooeffs[5]})
	}

	return (input)

}

func solveEquation(input Equation) int {

	a1 := float64(input.a1)
	a2 := float64(input.a2)
	b1 := float64(input.b1)
	b2 := float64(input.b2)
	c1 := float64(-1 * input.c1)
	c2 := float64(-1 * input.c2)

	a1a2 := a1 / a2
	b1b2 := b1 / b2
	c1c2 := c1 / c2

	// when the lines are parallel
	if a1a2 == b1b2 && b1b2 != c1c2 {
		return 0
	}

	// when the soln is consistent and independent
	if a1a2 != b1b2 {
		X := ((b1 * c2) - (b2 * c1)) / ((a1 * b2) - (a2 * b1))
		Y := ((c1 * a2) - (c2 * a1)) / ((a1 * b2) - (a2 * b1))

		if X < 0 || Y < 0 {
			return 0
		}

		//checking if there are no decimal places
		if (X/(float64(int(X))) == 1) && (Y/(float64(int(Y))) == 1) {
			return (3 * int(X)) + int(Y)
		}

		return 0

	} else {
		fmt.Println(input)
		fmt.Println("Same line")
		return 0
	}
}
