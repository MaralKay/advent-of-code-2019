package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ---------- NOTES ----------
// list of integers used for the initial state of the computer's memory
// opcodes mark the beginning of an instruction = opcode + parameters
// instruction pointer initially 0
// pointer increases (moves) by the number of values in the instruction (in our case 4 so far)

// find L(1) and L(2) => L(0) = 19690720
// L(1) : noun
// L(2) : verb
// calculate 100 * noun + verb

func main() {
	file, err := os.Open("puzzle2_input.txt")

	if err != nil {
		log.Fatal("Failed to open file.")
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var numbers = []int{}

		for scanner.Scan() {
			oneLine := scanner.Text()
			stringNumbers := strings.Split(oneLine, ",")
			for _, n := range stringNumbers {
				n, _ := strconv.Atoi(n)
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, n)
			}
		}
		file.Close()

		findInputForOutput(19690720, numbers)
	}
}

func findInputForOutput(wantedOutput int, inputArray []int) {
	var output int
loop:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			output = computer(inputArray, noun, verb)
			if output == wantedOutput {
				fmt.Println("100 * noun + verb = ", 100*noun+verb)
				break loop
			}
		}
	}
}

func computer(numbers []int, noun int, verb int) int {
	memory := make([]int, len(numbers))
	copy(memory, numbers)
	memory[1] = noun
	memory[2] = verb
	for i := 0; i < len(memory); i += 4 {
		if memory[i] != 99 {
			a := memory[memory[i+1]]
			b := memory[memory[i+2]]
			if memory[i] == 1 {
				memory[memory[i+3]] = a + b
			} else if memory[i] == 2 {
				memory[memory[i+3]] = a * b
			}
		} else {
			return memory[0]
		}
	}
	return memory[0]
}
