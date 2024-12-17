package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result int
	var tolerableResult int
	var failed int

	f, err := os.Open("input.txt")

	if err != nil {
		panic("Not able to open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		txt := scanner.Text()
		txtSlice := strings.Split(txt, " ")
		var intSlice []int

		for _, v := range txtSlice {
			i, _ := strconv.Atoi(v)
			intSlice = append(intSlice, i)
		}

		if isValidData(intSlice) {
			result++
		} else {
			res := isValidDataV2(intSlice)
			if res {
				tolerableResult++
			} else {
				failed++
			}
		}

	}

	fmt.Println(result)
	fmt.Println(tolerableResult)
	fmt.Println(failed)
	fmt.Println(result + tolerableResult)
}

func isValidData(a []int) bool {
	var slope int

	for i := 1; i < len(a); i++ {
		d := a[i] - a[i-1]

		if i == 1 {

			if d < 0 {
				slope = -1
			} else if d > 0 {
				slope = 1
			} else {
				return false
			}

		}

		if d*slope >= 1 && d*slope <= 3 {
			continue
		} else {

			return false
		}
	}
	return true
}

func isValidDataV2(a []int) bool {

	var considered bool = false
	fmt.Println("Up for consideration: ", a)

	for i := 1; i < len(a); i++ {

		if i == 1 {
			considered = isValidData(RemoveIndex(a, 0))
			if considered {
				return true
			}

		}
		d := a[i] - a[i-1]

		if i+1 < len(a) {
			d2 := a[i+1] - a[i]
			// if sign of slope channges
			if (d2 > 0) != (d > 0) {
				fmt.Println("Change  in direction of slope", i-1, i, i+1)
				considered = isValidData(RemoveIndex(a, i))
				if considered {

					return true
				}

				considered = isValidData(RemoveIndex(a, i-1))
				if considered {

					return true
				}
				considered = isValidData(RemoveIndex(a, i+1))
				if considered {

					return true
				}
				break
			}
		}

		if math.Abs(float64(d)) < 1 || math.Abs(float64(d)) > 3 {
			fmt.Println("Breaching threshold", i, i-1)
			considered = isValidData(RemoveIndex(a, i))
			if considered {

				return true
			}

			considered = isValidData(RemoveIndex(a, i-1))
			if considered {

				return true
			}
			break
		}
	}
	fmt.Println("Intolerable:", a)
	fmt.Println()
	return considered
}

func RemoveIndex(inSlice []int, index int) []int {
	aa := make([]int, len(inSlice))
	_ = copy(aa, inSlice)
	a := append(aa[:index], aa[index+1:]...)
	return a
}
