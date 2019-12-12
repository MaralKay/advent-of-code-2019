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
// list of comma separated integers: L
// start checking from position 0
// opcodes: 1, 2, 99
// 99 = program finished
// unknown opcode = error
// 1: and 3 next positions --> 1, a, b, c means write L(L(a)) + L(L(b)) to L(L(c))
// 2: and 3 next positions --> 2, a, b, c means write l(L(a)) * L(L(b)) to l(L(c))
// after processing opcode L(x), move to L(x+4)

func main() {
	file, err := os.Open("puzzle1_input.txt")

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

		numbers[1] = 12
		numbers[2] = 2

		result := computer(numbers)

		fmt.Println("Value at position 0: ", result[0])
	}
}

func computer(numbers []int) []int {
	for i := 0; i < len(numbers); i += 4 {
		a := numbers[numbers[i+1]]
		b := numbers[numbers[i+2]]
		if numbers[i] == 1 {
			numbers[numbers[i+3]] = a + b
		} else if numbers[i] == 2 {
			numbers[numbers[i+3]] = a * b
		} else if numbers[i] == 99 {
			return numbers
		} else {
			log.Fatal("Something went wrong.")
		}
	}
	return nil
}
