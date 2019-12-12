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
// of all individual fuel required for each module + their fuels + their fuels etc. till we reach a 0 or negative
// fuel value

func main() {
	file, err := os.Open("puzzle2_input.txt")

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
			fuel := calculateFuelString(mass)
			amountOfFuelRequired += fuel

			for {
				fuel = calculateFuelInt(fuel)
				if fuel <= 0 {
					break
				} else {
					amountOfFuelRequired += fuel
				}

			}
		}
		fmt.Println("amountOfFuelRequired:", amountOfFuelRequired)
	}
}

func calculateFuelString(x string) int {
	xAsFloat64, _ := strconv.ParseFloat(x, 64)
	return int(math.Floor(xAsFloat64/3)) - 2
}

func calculateFuelInt(x int) int {
	xAsFloat64 := float64(x)
	return int(math.Floor(xAsFloat64/3)) - 2
}
