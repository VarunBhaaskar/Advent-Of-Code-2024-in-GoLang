package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := "input.txt"
	input, i, j, fullMap := parse(path)

	var inl []int
	for in := 0; in < j; in++ {
		inl = append(inl, in)
	}
	fmt.Println("ii", inl)
	for k, line := range fullMap {
		fmt.Printf("%02d %v", k, line)
		fmt.Println()
	}

	var resultMap = make(map[[2]int]int)
	var resultMap2 = make(map[[2]int]int)
	for _, v := range input {
		o, o2 := checkAntiNodes(v, i, j)
		for _, v := range o {
			resultMap[v] += 1
		}

		for _, v := range o2 {
			resultMap2[v] += 1
		}
	}
	fmt.Println(resultMap)
	fmt.Println(resultMap2)
	fmt.Println("Part 1 result: ", len(resultMap))
	fmt.Println("Part 2 result: ", len(resultMap2))
}

func parse(path string) (map[string][][2]int, int, int, [][]string) {

	input := make(map[string][][2]int)
	var fullMap [][]string
	f, err := os.Open(path)

	if err != nil {
		panic("Cant open file")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	i_index := 0
	j_index := 0
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

		j_index = len(line_slice)
		i_index++
	}

	defer f.Close()

	return input, i_index, j_index, fullMap
}

func checkAntiNodes(input [][2]int, i, j int) ([][2]int, [][2]int) {
	var output [][2]int
	var output2 [][2]int
	n := len(input)
	for k := 0; k < n-1; k++ {
		for m := k + 1; m < n; m++ {
			a, b := formAntiNodes(input[k], input[m])

			if a[0] < i && a[1] < j && a[0] >= 0 && a[1] >= 0 {
				output = append(output, a)
			}
			if b[0] < i && b[1] < j && b[0] >= 0 && b[1] >= 0 {
				output = append(output, b)
			}

			antiNodesWithHarmonics := formAntiNodesWithHarmonics(input[k], input[m], i, j)
			output2 = append(output2, antiNodesWithHarmonics...)
		}
	}
	fmt.Println(output)
	return output, output2
}

func formAntiNodes(a [2]int, b [2]int) ([2]int, [2]int) {

	i_diff := b[0] - a[0]
	j_diff := a[1] - b[1]

	// if j_diff >= 0 {
	return [2]int{a[0] - i_diff, a[1] + j_diff}, [2]int{b[0] + i_diff, b[1] - j_diff}
	// } else {
	// 	return [2]int{a[0] - i_diff, a[1] + j_diff}, [2]int{b[0] + i_diff, b[1] - j_diff}
	// }

}

func formAntiNodesWithHarmonics(a [2]int, b [2]int, i, j int) [][2]int {
	// fmt.Println("Starting harmonis:", a, b, i, j)
	var allAntinodes [][2]int

	i_diff := b[0] - a[0]
	j_diff := a[1] - b[1]

	for {

		if a[0] < i && a[1] < j && a[0] >= 0 && a[1] >= 0 {
			allAntinodes = append(allAntinodes, a)
		} else {
			break
		}
		a = [2]int{a[0] - i_diff, a[1] + j_diff}

	}

	for {

		if b[0] < i && b[1] < j && b[0] >= 0 && b[1] >= 0 {
			allAntinodes = append(allAntinodes, b)
		} else {
			break
		}
		b = [2]int{b[0] + i_diff, b[1] - j_diff}
	}

	return allAntinodes
}
