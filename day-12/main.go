package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

func main() {
	path := "input.txt"
	mapInput, input := parse(path)

	// mapInput2 := make(map[string][][2]int)

	// for k, v := range mapInput {
	// 	mapInput2[k] = v
	// }

	// result := findBoundaries(mapInput, input)
	// fmt.Println("part 1 result: ", result)
	// fmt.Println(mapInput2)
	result2 := findBoundaries2(mapInput, input)
	fmt.Println("part 2 result: ", result2)
}

func parse(path string) (map[string][][2]int, [][]string) {

	input := make(map[string][][2]int)
	var fullMap [][]string
	f, err := os.Open(path)

	if err != nil {
		panic("Cant open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	i_index := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_slice := strings.Split(line, "")
		// fmt.Println(line_slice)
		fullMap = append(fullMap, line_slice)
		for j, v := range line_slice {
			if v != "." {
				input[v] = append(input[v], [2]int{i_index, j})

			}
		}
		i_index++
	}

	return input, fullMap
}

func findBoundaries(mapInput map[string][][2]int, input [][]string) int {
	// var result [][]Point
	var finalRes int

start:
	for k, v := range mapInput {
		for _, x := range v {
			result := make(map[[2]int]int)
			points := nextAvailablePoints(input, x[0], x[1], result)
			// fmt.Println(k, points)

			perimeter := 0

			// fmt.Println("Before", mapInput[k])
			for pk, pv := range points {
				perimeter += pv
				in := slices.Index(mapInput[k], pk)
				mapInput[k] = deleteElement(mapInput[k], in)
			}
			// fmt.Println("After", mapInput[k])

			// fmt.Println(mapInput)
			finalRes += perimeter * len(points)
			// fmt.Println("ans is ", perimeter, finalRes)
			goto start
		}
	}

	return finalRes
}

func nextAvailablePoints(input [][]string, i, j int, result map[[2]int]int) map[[2]int]int {

	val := input[i][j]

	if _, ok := result[[2]int{i, j}]; ok {
		return result
	} else {
		result[[2]int{i, j}] = 4
	}

	maxI := len(input)
	maxJ := len(input)

	if i-1 >= 0 && input[i-1][j] == val {
		result[[2]int{i, j}]--
		result = nextAvailablePoints(input, i-1, j, result)
	}
	if j-1 >= 0 && input[i][j-1] == val {
		result[[2]int{i, j}]--
		result = nextAvailablePoints(input, i, j-1, result)
	}
	if i+1 < maxI && input[i+1][j] == val {
		result[[2]int{i, j}]--
		result = nextAvailablePoints(input, i+1, j, result)
	}
	if j+1 < maxJ && input[i][j+1] == val {
		result[[2]int{i, j}]--
		result = nextAvailablePoints(input, i, j+1, result)
	}

	return result

}

func deleteElement[T any](input []T, index int) []T {
	if index < 0 || index >= len(input) {
		return input
	}
	if index == len(input)-1 {
		return input[:index]
	}

	if index == 0 {
		return input[1:]
	}
	return append(input[:index], input[index+1:]...)
}

func findBoundaries2(mapInput map[string][][2]int, input [][]string) int {
	// var result [][]Point
	var finalRes int
start:
	for k, v := range mapInput {
		for _, x := range v { // this loop is just for accessing a random key from map input. will not run once no elements remain. and will not reach start:
			result := make(map[[2]int][]string)
			points := nextAvailablePoints2(input, x[0], x[1], result)
			// fmt.Println(k)

			// perimeter := make(map[string]int)

			// fmt.Println("Before", mapInput[k])
			for pk := range points {

				// for _, y := range pv {
				// 	perimeter[y]++
				// }

				in := slices.Index(mapInput[k], pk)
				mapInput[k] = deleteElement(mapInput[k], in)
			}
			// fmt.Println("After", mapInput[k])

			// fmt.Println(mapInput)
			// sides := calculateSides(points, len(input), len(input[0]))
			sides := calculateSides(points, len(input), len(input[0]))
			finalRes += sides * len(points)
			// fmt.Println(k, "ans is ", sides, len(points))
			goto start
		}
	}

	return finalRes
}

func nextAvailablePoints2(input [][]string, i, j int, result map[[2]int][]string) map[[2]int][]string {
	// encoding edges to find unique at the end
	val := input[i][j]

	if _, ok := result[[2]int{i, j}]; ok {
		return result
	}

	result[[2]int{i, j}] = []string{}

	maxI := len(input)
	maxJ := len(input)

	if i-1 >= 0 && input[i-1][j] == val {
		result = nextAvailablePoints2(input, i-1, j, result)
	} else {
		result[[2]int{i, j}] = append(result[[2]int{i, j}], fmt.Sprintf("i-(%d)-(%d)", i, i-1))
	}

	if j-1 >= 0 && input[i][j-1] == val {

		result = nextAvailablePoints2(input, i, j-1, result)
	} else {
		result[[2]int{i, j}] = append(result[[2]int{i, j}], fmt.Sprintf("j-(%d)-(%d)", j, j-1))
	}

	if i+1 < maxI && input[i+1][j] == val {
		result = nextAvailablePoints2(input, i+1, j, result)
	} else {
		result[[2]int{i, j}] = append(result[[2]int{i, j}], fmt.Sprintf("i-(%d)-(%d)", i, i+1))
	}

	if j+1 < maxJ && input[i][j+1] == val {

		result = nextAvailablePoints2(input, i, j+1, result)
	} else {
		result[[2]int{i, j}] = append(result[[2]int{i, j}], fmt.Sprintf("j-(%d)-(%d)", j, j+1))
	}
	return result

}

func search2(arr []string, k string) (int, bool) {
	for i, v := range arr {
		if v == k {
			return i, true
		}
	}

	return 0, false
}

func calculateSides(points map[[2]int][]string, i_max, j_max int) int {

	// looks weird and complex for sure. What i am doing is, when i find a edge
	// encoded as i-{n1}-{n2} is an horizontal edge between rows n1 -> n2
	// encoded as j-{n1}-{n2} is a vertical edge between columns n1 -> n2
	// when i find an edge within a fence, i start iterating until the boundaries of the input matrix.
	// if the flow breaks, all the edges before the flow break constitute an side
	// this will account for direction j-2-3 being diiferent from j-3-2
	// and also account for break in continuity of the same edge to be counted as 2 sides
	// there could be 2 j-2-3 but disconnected
	// iterating from the (X, Y) coord of point associated to edge  until (i_max, j_max)
	// Time Complexity : O(F)
	// fmt.Println(points)
	var sides []string

	keys := make([][2]int, 0, len(points))
	for k := range points {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i2, j2 int) bool {
		if keys[i2][0] == keys[j2][0] {
			return keys[i2][1] < keys[j2][1]
		}
		return keys[i2][0] < keys[j2][0]
	})

	// fmt.Println("Keys ", keys)

	for k := 0; k < len(keys); {
		x := keys[k]
		// fmt.Println("XX is", k, x, points[x])
		if len(points[x]) > 0 {
			side := points[x][0]
			// fmt.Println("Side is ", side)
			sides = append(sides, side)
			points[x] = deleteElement(points[x], 0)
			if string(side[0]) == "i" {
				for j := x[1] + 1; j < j_max; j++ {
					if peri, ok := points[[2]int{x[0], j}]; ok && len(points[[2]int{x[0], j}]) > 0 {
						// fmt.Println("Peri", [2]int{x[0], j}, peri)
						if index, ok := search2(peri, side); ok {
							points[[2]int{x[0], j}] = deleteElement(points[[2]int{x[0], j}], index)
						} else {
							break
						}
					} else {
						break
					}
				}
			} else {
				for i := x[0] + 1; i < i_max; i++ {
					if peri, ok := points[[2]int{i, x[1]}]; ok && len(points[[2]int{i, x[1]}]) > 0 {
						// fmt.Println("Peri", [2]int{i, x[1]}, peri)
						if index, ok := search2(peri, side); ok {
							points[[2]int{i, x[1]}] = deleteElement(points[[2]int{i, x[1]}], index)
						} else {
							break
						}
					} else {
						break
					}
				}
			}

		} else {
			k++
		}
	}
	// fmt.Println("Sides ", sides)
	return len(sides)
}
