package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// we need to find the value of amountOfFuelRequired which is the sum
// of all individual fuel required for each module

// for m in masses:
// totalAmountOfFuel += roundDown(mass/3) - 2
// end

func main() {
	file, err := os.Open("puzzle1_input.txt")

	if err != nil {
		log.Fatal("Failed to open file.")
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var masses []string

		for scanner.Scan() {
			masses = append(masses, scanner.Text())
		}
		file.Close()

		amountOfFuelRequired := 0
		for _, mass := range masses {
			massFloat64, _ := strconv.ParseFloat(mass, 64)
			amountOfFuelRequired += int(math.Floor(massFloat64/3)) - 2
		}
		fmt.Println("amountOfFuelRequired:", amountOfFuelRequired)
	}
}
