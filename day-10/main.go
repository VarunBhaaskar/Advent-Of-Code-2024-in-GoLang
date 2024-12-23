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
	input, zeros := parse(path)

	var result int
	var result2 int
	for _, z := range zeros {
		result += findPath(input, z, make(map[[2]int]int))

	}

	for _, z := range zeros {
		result2 += findPath2(input, z)

	}

	fmt.Println("Part 1 result: ", result)
	fmt.Println("Part 2 result: ", result2)
}

func parse(path string) ([][]int, [][2]int) {
	f, err := os.Open(path)
	var zeroes [][2]int
	var input [][]int

	if err != nil {
		panic("Cant open file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	index_i := 0
	for scanner.Scan() {
		line := scanner.Text()

		line_slc := strings.Split(line, "")
		var tmp []int

		for j, x := range line_slc {
			v, err := strconv.Atoi(x)

			if err != nil {
				v = -1
			}

			if v == 0 {
				zeroes = append(zeroes, [2]int{index_i, j})
			}

			tmp = append(tmp, v)
		}
		input = append(input, tmp)
		index_i++
	}

	fmt.Println(zeroes)
	fmt.Println(input)

	return input, zeroes
}

func findPath(input [][]int, pos [2]int, visitedPoints map[[2]int]int) int {
	max_i := len(input)
	max_j := len(input[0])
	// fmt.Println(pos, input[pos[0]][pos[1]])
	var result int

	value := input[pos[0]][pos[1]]

	if _, ok := visitedPoints[pos]; ok {
		return 0
	} else {
		visitedPoints[pos] = 1
	}

	if input[pos[0]][pos[1]] == 9 {
		// fmt.Println("Returned after rwching", pos)
		return 1
	}

	if pos[0]-1 >= 0 && input[pos[0]-1][pos[1]] == value+1 {
		// fmt.Println("going up from", pos, "to")
		result += findPath(input, [2]int{pos[0] - 1, pos[1]}, visitedPoints)
	}

	if pos[1]-1 >= 0 && input[pos[0]][pos[1]-1] == value+1 {
		// fmt.Println(" going left from", pos)
		result += findPath(input, [2]int{pos[0], pos[1] - 1}, visitedPoints)
	}

	if pos[0]+1 < max_i && input[pos[0]+1][pos[1]] == value+1 {
		// fmt.Println("going down from", pos)
		result += findPath(input, [2]int{pos[0] + 1, pos[1]}, visitedPoints)
	}

	if pos[1]+1 < max_j && input[pos[0]][pos[1]+1] == value+1 {
		// fmt.Println(" going right from", pos)
		result += findPath(input, [2]int{pos[0], pos[1] + 1}, visitedPoints)
	}

	return result
}

func findPath2(input [][]int, pos [2]int) int {
	max_i := len(input)
	max_j := len(input[0])
	// fmt.Println(pos, input[pos[0]][pos[1]])
	var result int

	value := input[pos[0]][pos[1]]

	if input[pos[0]][pos[1]] == 9 {
		// fmt.Println("Returned after rwching", pos)
		return 1
	}

	if pos[0]-1 >= 0 && input[pos[0]-1][pos[1]] == value+1 {
		// fmt.Println("going up from", pos, "to")
		result += findPath2(input, [2]int{pos[0] - 1, pos[1]})
	}

	if pos[1]-1 >= 0 && input[pos[0]][pos[1]-1] == value+1 {
		// fmt.Println(" going left from", pos)
		result += findPath2(input, [2]int{pos[0], pos[1] - 1})
	}

	if pos[0]+1 < max_i && input[pos[0]+1][pos[1]] == value+1 {
		// fmt.Println("going down from", pos)
		result += findPath2(input, [2]int{pos[0] + 1, pos[1]})
	}

	if pos[1]+1 < max_j && input[pos[0]][pos[1]+1] == value+1 {
		// fmt.Println(" going right from", pos)
		result += findPath2(input, [2]int{pos[0], pos[1] + 1})
	}

	return result
}
