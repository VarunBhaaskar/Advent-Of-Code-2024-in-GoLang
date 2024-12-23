package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Chunks struct {
	value string
	size  int
}

func main() {

	path := "input.txt"
	input, chunkInput := parse(path)

	swappedInput := fillGaps(input)
	optimized := fillGapsOptimum(chunkInput)

	var result int
	var result2 int

	for i, v := range swappedInput {
		if v == "." {
			break
		}

		vi, _ := strconv.Atoi(v)
		result = result + (i * vi)
	}
	for i, v := range optimized {
		if v == "." {
			continue
		}

		vi, _ := strconv.Atoi(v)
		result2 = result2 + (i * vi)
	}
	fmt.Println("Part 1 result - ", result)
	fmt.Println("Part 2 result - ", result2)
}

func parse(path string) ([]string, []Chunks) {
	var input []string
	var chunkInput []Chunks

	f, err := os.Open(path)

	if err != nil {
		panic("Cant open input file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanBytes)

	index := 0
	file_index := 0
	for scanner.Scan() {
		a := scanner.Text()
		i, _ := strconv.Atoi(a)
		if index%2 == 0 {
			chunkInput = append(chunkInput, Chunks{strconv.Itoa(file_index), i})
			for j := 0; j < i; j++ {
				input = append(input, strconv.Itoa(file_index))

			}
			file_index++

		} else {
			chunkInput = append(chunkInput, Chunks{".", i})
			for j := 0; j < i; j++ {
				input = append(input, ".")
			}
		}

		index++
	}
	// fmt.Println(chunkInput)
	return input, chunkInput

}

func fillGaps(input []string) []string {
	n := len(input)

	indexToSwapWith := n - 1

	for i := 0; i < n; i++ {
		if input[i] != "." {
			continue
		} else {
			for {
				if input[indexToSwapWith] != "." {
					break
				} else {
					indexToSwapWith--
				}
			}

			if indexToSwapWith <= i {
				break
			}
			input[indexToSwapWith], input[i] = input[i], input[indexToSwapWith]
			// fmt.Println(input)
			indexToSwapWith--
		}
	}

	return input
}

func fillGapsOptimum(input []Chunks) []string {
	n := len(input)
	indexToSwapWith := 0

	for i := n - 1; i >= 0; i-- {

		if input[i].value == "." {
			continue
		}

		if indexToSwapWith >= i {
			break
		}

		//finding first available free space
		for j := indexToSwapWith; j < i; j++ {

			if input[j].value != "." || input[j].size < input[i].size {
				continue
			}
			tmpInput := make([]Chunks, len(input))
			copy(tmpInput, input)

			// fmt.Println("Starting:", i, indexToSwapWith, input[i], input[j])
			diff := input[j].size - input[i].size

			inter := append(tmpInput[:j], input[i])

			input[i].value = "."
			// fmt.Println("First", inter, " - ", input, " - ", Chunks{".", diff})
			inter = append(inter, Chunks{".", diff})
			// fmt.Println("Second", inter, " - ", input)

			inter = append(inter, input[j+1:]...)
			// fmt.Println("Third", inter, " - ", input)
			// fmt.Println("After Swapping: ", inter)
			input = inter

			break
		}
	}

	var sliceInput []string

	for _, v := range input {

		for i := 0; i < v.size; i++ {
			sliceInput = append(sliceInput, v.value)
		}
	}
	// fmt.Println(sliceInput)

	return sliceInput

}
