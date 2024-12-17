package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {

	start := time.Now()
	path := "input.txt"
	var result int
	var result2 int

	input := parse(path)
	// fmt.Println(input)

	for i, row := range input {
		for j, v := range row {
			if v == "X" {

				if checkBackward(i, j, input) {
					result++
				}

				if checkForward(i, j, input) {
					result++
				}

				if checkToptoBottom(i, j, input) {
					result++
				}

				if checkBottomToTop(i, j, input) {
					result++
				}

				if checkleftDianol(i, j, input) {
					result++
				}

				if checkRightDiagnol(i, j, input) {
					result++
				}
				if checkleftDianolReverse(i, j, input) {
					result++
				}
				if checkRightDiagnolReverse(i, j, input) {
					result++
				}

				// fmt.Println(i, j, result)
			}

			chan2 := make(chan bool, 4)

			if v == "M" {
				wg.Add(1)
				go forwardMAS(i, j, input, chan2)
				wg.Add(1)
				go backwardMAS(i, j, input, chan2)
				wg.Add(1)
				go forwardMMSS(i, j, input, chan2)
				wg.Add(1)
				go BackwardMMSS(i, j, input, chan2)
			}
			go func() {
				wg.Wait()
				close(chan2)
			}()
			for boolv := range chan2 {
				if boolv {
					result2++
				}
			}
		}

	}
	fmt.Println("Part-1: ", result)
	fmt.Println("Part-2: ", result2)
	end := time.Since(start)
	fmt.Println("Took: ", end)
}

func checkBackward(i, j int, input [][]string) bool {
	if j-3 < 0 {
		return false
	}
	// fmt.Println(strings.Join(input[i][j-3:j+1], ""))
	if input[i][j-1] == "M" && input[i][j-2] == "A" && input[i][j-3] == "S" {
		return true
	} else {
		return false
	}
}

func checkForward(i, j int, input [][]string) bool {

	if j+3 >= len(input[i]) {
		return false
	}
	// fmt.Println(strings.Join(input[i][j:j+4], ""))
	if strings.Join(input[i][j:j+4], "") == "XMAS" {
		return true
	} else {
		return false
	}

}

func checkToptoBottom(i, j int, input [][]string) bool {
	if i+3 >= len(input) {
		return false
	}

	if input[i+1][j] == "M" && input[i+2][j] == "A" && input[i+3][j] == "S" {
		return true
	} else {
		return false
	}

}

func checkBottomToTop(i, j int, input [][]string) bool {
	if i-3 < 0 {
		return false
	}

	if input[i-1][j] == "M" && input[i-2][j] == "A" && input[i-3][j] == "S" {
		return true
	} else {
		return false
	}
}

func checkleftDianol(i, j int, input [][]string) bool {
	if i+3 >= len(input) || j-3 < 0 {
		return false
	}
	if input[i+1][j-1] == "M" && input[i+2][j-2] == "A" && input[i+3][j-3] == "S" {
		return true
	} else {
		return false
	}
}

func checkRightDiagnol(i, j int, input [][]string) bool {
	if i+3 >= len(input) || j+3 >= len(input[i]) {
		return false
	}
	if input[i+1][j+1] == "M" && input[i+2][j+2] == "A" && input[i+3][j+3] == "S" {
		return true
	} else {
		return false
	}
}

func checkleftDianolReverse(i, j int, input [][]string) bool {
	if i-3 < 0 || j-3 < 0 {
		return false
	}
	if input[i-1][j-1] == "M" && input[i-2][j-2] == "A" && input[i-3][j-3] == "S" {
		return true
	} else {
		return false
	}
}

func checkRightDiagnolReverse(i, j int, input [][]string) bool {
	if i-3 < 0 || j+3 >= len(input[i]) {
		return false
	}
	if input[i-1][j+1] == "M" && input[i-2][j+2] == "A" && input[i-3][j+3] == "S" {
		return true
	} else {
		return false
	}
}

func parse(path string) (input [][]string) {

	f, err := os.Open(path)

	if err != nil {
		panic("Cant read input")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		line_slice := strings.Split(line, "")
		input = append(input, line_slice)
	}

	return

}

func forwardMAS(i, j int, input [][]string, channel chan bool) {
	defer wg.Done()

	if j+2 >= len(input[i]) || i+2 >= len(input) {
		channel <- false

		return
	}

	if input[i][j+2] == "S" && input[i+1][j+1] == "A" && input[i+2][j] == "M" && input[i+2][j+2] == "S" {
		// fmt.Println(i, j, input[i][j+2], input[i+1][j+1], input[i+2][j], input[i+2][j+2])
		channel <- true
	} else {
		channel <- false
	}

}

func backwardMAS(i, j int, input [][]string, channel chan bool) {
	defer wg.Done()

	if j-2 < 0 || i-2 < 0 {
		channel <- false

		return
	}

	if input[i][j-2] == "S" && input[i-1][j-1] == "A" && input[i-2][j] == "M" && input[i-2][j-2] == "S" {
		// fmt.Println(i, j, input[i][j-2], input[i-1][j-1], input[i-2][j], input[i-2][j-2])
		channel <- true
	} else {
		channel <- false
	}

}

func forwardMMSS(i, j int, input [][]string, channel chan bool) {
	defer wg.Done()

	if j+2 >= len(input[i]) || i+2 >= len(input) {
		channel <- false

		return
	}

	if input[i][j+2] == "M" && input[i+1][j+1] == "A" && input[i+2][j] == "S" && input[i+2][j+2] == "S" {
		// fmt.Println(i, j, input[i][j+2], input[i+1][j+1], input[i+2][j], input[i+2][j+2])
		channel <- true
	} else {
		channel <- false
	}

}

func BackwardMMSS(i, j int, input [][]string, channel chan bool) {
	defer wg.Done()
	if j-2 < 0 || i-2 < 0 {
		channel <- false

		return
	}

	if input[i][j-2] == "M" && input[i-1][j-1] == "A" && input[i-2][j] == "S" && input[i-2][j-2] == "S" {
		// fmt.Println(i, j, input[i][j-2], input[i-1][j-1], input[i-2][j], input[i-2][j-2])
		channel <- true
	} else {
		channel <- false
	}

}
