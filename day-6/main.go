package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PositionEncoding struct {
	i         int
	j         int
	direction string
}

func main() {
	path := "input.txt"
	input, i, j := parse(path)
	fmt.Println(input)
	fmt.Println("Starting pos: ", i, j)

	result1, traversedWithDirection := move(input, i, j)
	fmt.Println(traversedWithDirection)

	var result2 int

	for k := range result1 {
		input[k[0]][k[1]] = "#"

		if moveLoopDetector(input, i, j) {
			result2++
		}

		input[k[0]][k[1]] = "."
	}

	fmt.Println("part 1 result: ", len(result1))
	fmt.Println("Part 2 result: ", result2)
}

func parse(path string) ([][]string, int, int) {

	var input [][]string
	var found bool = false
	var i, j int
	f, err := os.Open(path)

	if err != nil {
		panic("Cant open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "")
		input = append(input, line)

		if !found {
			// fmt.Println(i, line)
			if v, ok := search2(line, string(rune(94))); ok {
				j = v
				found = true
			} else {
				i++
			}
		}
	}

	return input, i, j
}

func search2(arr []string, k string) (int, bool) {
	for i, v := range arr {
		if v == k {
			return i, true
		}
	}

	return 0, false
}

func newDirection(old string) string {
	//  F -> R -> B -> L

	switch old {
	case "F":
		return "R"
	case "R":
		return "B"
	case "B":
		return "L"
	case "L":
		return "F"
	default:
		return old
	}
}

func move(input [][]string, i, j int) (map[[2]int]int, map[PositionEncoding]int) {
	traversed := make(map[[2]int]int)
	traversed[[2]int{i, j}] = 1
	n1 := len(input)
	n2 := len(input[0])
	direction := "F"

	traversedWithDirection := make(map[PositionEncoding]int)

	for {
		switch direction {
		case "F":
			if i-1 < 0 {
				return traversed, traversedWithDirection
			}

			if input[i-1][j] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				i = i - 1

				if mv, ok := traversed[[2]int{i, j}]; ok {
					traversed[[2]int{i, j}] = mv + 1
				} else {
					traversed[[2]int{i, j}] = 1
				}
				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		case "R":
			if j+1 >= n2 {
				return traversed, traversedWithDirection
			}
			if input[i][j+1] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				j = j + 1

				if mv, ok := traversed[[2]int{i, j}]; ok {
					traversed[[2]int{i, j}] = mv + 1
				} else {
					traversed[[2]int{i, j}] = 1
				}
				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		case "B":
			if i+1 >= n1 {
				return traversed, traversedWithDirection
			}
			if input[i+1][j] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				i = i + 1

				if mv, ok := traversed[[2]int{i, j}]; ok {
					traversed[[2]int{i, j}] = mv + 1
				} else {
					traversed[[2]int{i, j}] = 1
				}
				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		case "L":
			if j-1 < 0 {
				return traversed, traversedWithDirection
			}
			if input[i][j-1] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				j = j - 1

				if mv, ok := traversed[[2]int{i, j}]; ok {
					traversed[[2]int{i, j}] = mv + 1
				} else {
					traversed[[2]int{i, j}] = 1
				}
				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		default:
			fmt.Println("Somehow reached default")
			return traversed, traversedWithDirection

		}
	}

}

func loopHim(traversedWithDirection map[PositionEncoding]int, input [][]string) int {
	var result int
	fmt.Println(traversedWithDirection)

	n1 := len(input)
	n2 := len(input[0])

	fmt.Println(n1, n2)

	for pos := range traversedWithDirection {

		if pos.direction == "F" {
			// fmt.Println(input[pos.i+1][pos.j+1:])
			if j, ok := search2(input[pos.i+1][pos.j+1:], "#"); !ok {
				continue
			} else {
				i_3 := -1
				for iter_i := pos.i + 1; iter_i < n1; iter_i++ {
					if input[iter_i][j-1] == "#" {
						i_3 = iter_i - 1
						break
					}
				}
				if i_3 >= 0 {
					if pos.j-1 >= 0 && input[i_3][pos.j-1] == "#" {
						result++
						fmt.Println(pos)
					}
				} else {
					continue
				}
			}

		}

		if pos.direction == "B" {
			// fmt.Println(pos)
			// fmt.Println(input[pos.i-1][0:pos.j])

			if j, ok := search2(input[pos.i-1][0:pos.j], "#"); !ok {
				// fmt.Println("Search failed", j, ok)
				continue
			} else {
				// fmt.Println(j)
				i_3 := -1 // i coordinate of 3rd point in the loop assuming our obstacle is the first
				for iter_i := pos.i - 1; iter_i >= 0; iter_i-- {
					// fmt.Println(iter_i, j+1)
					if input[iter_i][j+1] == "#" {
						i_3 = iter_i + 1
						break
					}
				}
				// fmt.Println(i_3)
				if i_3 >= 0 {
					if pos.j+1 < n2 && input[i_3][pos.j+1] == "#" {
						result++
						fmt.Println("Success(B):", pos)
					}
				} else {
					continue
				}
			}

		}

		if pos.direction == "R" {
			i_2 := -1
			for iter_i := pos.i + 1; iter_i < n1; iter_i++ {
				if input[iter_i][pos.j-1] == "#" {
					i_2 = iter_i - 1
					break
				}
			}

			if i_2 < 0 {
				continue
			}

			if j_3, ok := search2(input[i_2][:pos.j-1], "#"); !ok {
				continue
			} else {
				if j_3+1 < n2 && input[pos.i-1][j_3+1] == "#" {
					result++
					fmt.Println("Success(R):", pos)
				}
			}
		}

	}

	return result
}

func moveLoopDetector(input [][]string, i, j int) bool {

	n1 := len(input)
	n2 := len(input[0])
	direction := "F"

	traversedWithDirection := make(map[PositionEncoding]int)

	startingPOS := PositionEncoding{i, j, direction}

	traversedWithDirection[startingPOS] = 1

	for {
		switch direction {
		case "F":
			if i-1 < 0 {
				return false
			}

			if input[i-1][j] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				i = i - 1

				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
					return true
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		case "R":
			if j+1 >= n2 {
				return false
			}
			if input[i][j+1] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				j = j + 1

				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
					return true
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		case "B":
			if i+1 >= n1 {
				return false
			}
			if input[i+1][j] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				i = i + 1

				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
					return true
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		case "L":
			if j-1 < 0 {
				return false
			}
			if input[i][j-1] == "#" {
				direction = newDirection(direction)
				continue
			} else {
				j = j - 1
				position := PositionEncoding{i, j, direction}
				if md, ok := traversedWithDirection[position]; ok {
					traversedWithDirection[position] = md + 1
					return true
				} else {
					traversedWithDirection[position] = 1
				}
				continue
			}
		default:
			fmt.Println("Somehow reached default")
			return false

		}
	}

}
