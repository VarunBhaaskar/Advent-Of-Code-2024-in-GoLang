package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	sample()
	part1()
	part2()
}

type Velocity struct {
	start_pos  [2]int
	i_velocity int // velocity along a column
	j_velocity int // veocity along a row
}

func sample() {
	path := "sample.txt"

	// var input [7][11][]Velocity
	// fmt.Println(input)

	i := 7
	j := 11
	time := 100
	var resultMat [7][11]int
	var resMap = make(map[[2]int]int)

	for v := range parse(path) {
		// fmt.Println(v)

		ans := move(v, i, j, time)
		resultMat[ans[0]][ans[1]]++
		resMap[ans]++
	}

	fmt.Println("Sample result: ", makeQuadrant(resMap, i, j))

}

func part1() {
	path := "input.txt"

	// var input [7][11][]Velocity
	// fmt.Println(input)

	i := 103
	j := 101
	time := 100
	var resultMat [103][101]int
	var resMap = make(map[[2]int]int)

	for v := range parse(path) {
		// fmt.Println(v)

		ans := move(v, i, j, time)
		resultMat[ans[0]][ans[1]]++
		resMap[ans]++
	}

	fmt.Println("Part 1 result: ", makeQuadrant(resMap, i, j))

}

func part2() {
	path := "input.txt"

	// var input [7][11][]Velocity
	// fmt.Println(input)
	os.Remove("results.txt")
	i := 103
	j := 101

	var point []Velocity
	for v := range parse(path) {
		// fmt.Println(v)

		point = append(point, v)
	}

	for x := 0; x < i*j; x++ {
		// fmt.Println(x, " Seconds have passed")
		for y := 0; y < len(point); y++ {
			point[y].start_pos = move(point[y], i, j, 1)
		}
		displayMap(point, x)
	}

}

func parse(path string) iter.Seq[Velocity] {

	return func(yield func(Velocity) bool) {
		f, err := os.Open(path)

		if err != nil {
			panic("Cannt open file")
		}

		defer f.Close()

		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)

		re := regexp.MustCompile(`p=(\d{1,}),(\d{1,}) v=([-]{0,1}\d{1,}),([-]{0,1}\d{1,})`)

		for scanner.Scan() {
			line := scanner.Text()
			values := re.FindStringSubmatch(line)
			var values_int [4]int

			for i := 1; i < 5; i++ {
				values_int[i-1], _ = strconv.Atoi(values[i])
			}
			// fmt.Println(values[1:])
			//swapping velocity positions so that i velocity and j velocity matches
			// swapping p as well so that 1t  element describes row and second represents column
			res := Velocity{[2]int{values_int[1], values_int[0]}, values_int[3], values_int[2]}

			if !yield(res) {
				return
			}
		}
	}

}

func move(initialPos Velocity, i, j, time int) [2]int {

	pos := initialPos.start_pos

	iv := initialPos.i_velocity
	jv := initialPos.j_velocity

	for x := 0; x < time; x++ {
		pos[0] = pos[0] + iv

		if pos[0] < 0 {
			pos[0] = i + pos[0]
		} else if pos[0] >= i {
			pos[0] = pos[0] - i
		}

		pos[1] = pos[1] + jv

		if pos[1] < 0 {
			pos[1] = j + pos[1]
		} else if pos[1] >= j {
			pos[1] = pos[1] - j
		}
		// fmt.Println(pos)
	}

	return pos
}

func makeQuadrant(res map[[2]int]int, i, j int) int {
	var quad1, quad2, quad3, quad4 int
	// fmt.Println(i/2, j/2)

	i2 := i / 2
	j2 := j / 2
	for k, v := range res {
		// fmt.Println(k, v)
		if k[0] == i2 {
			continue
		} else if k[1] == j2 {
			continue
		}

		if k[0] < i2 && k[1] > j2 {
			quad1 += v
		} else if k[0] < i2 && k[1] < j2 {
			quad2 += v
		} else if k[0] > i2 && k[1] < j2 {
			quad3 += v
		} else {
			quad4 += v
		}

	}

	return quad1 * quad2 * quad3 * quad4
}

func displayMap(input []Velocity, x int) {
	var matrix [103][101]string

	for i := 0; i < 103; i++ {
		for j := 0; j < 101; j++ {
			matrix[i][j] = "."
		}
	}

	for _, v := range input {
		matrix[v.start_pos[0]][v.start_pos[1]] = "#"
	}

	file, err := os.OpenFile("results.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("failed creating file:")
	}
	defer file.Close()
	datawriter := bufio.NewWriter(file)
	datawriter.WriteString("\n" + "After " + strconv.Itoa(x+1) + " Seconds" + "\n")
	for _, v := range matrix {

		datawriter.WriteString(strings.Join(v[:], "") + "\n")
	}

	datawriter.Flush()
}
