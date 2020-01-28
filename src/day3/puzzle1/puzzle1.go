package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Manhattan distance: d(p,q) = ||p-q||1 = Si...n|p-q|i where p,q vectors
// e.g.: distance between (p1, p2) and (q1, q2) is |p1-q1|+|p2-q2|

// we need to translate the wire information we have into some arithmetic value as positions
// we will do that with a pair of numbers, one for axis x, one for axis y

// Point is a pair of coordinates on the cartesian axes
type Point struct {
	x int
	y int
}

// Wire is described by an array of Points
type Wire struct {
	coordinates []Point
}

func main() {
	var wires [][]string
	file, err := os.Open("puzzle1_input.txt")

	if err != nil {
		log.Fatal("Failed to open file.")
	} else {
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			wire := scanner.Text()
			wireDirections := strings.Split(wire, ",")
			wires = append(wires, wireDirections)
		}

	}
	file.Close()

	// right gives +1, left gives -1
	var rightLeft []int
	// up gives +1, down gives -1
	var upDown []int

	var wireCoordinates []Wire
	for i := 0; i < len(wires); i++ {
		var coordinates []Point
		rightLeft = append(rightLeft, 0)
		upDown = append(upDown, 0)
		for j := 0; j < len(wires[i]); j++ {
			if strings.Contains(wires[i][j], "R") {
				r, _ := strconv.Atoi(strings.Replace(wires[i][j], "R", "", 1))
				rightLeft[i] += r
			} else if strings.Contains(wires[i][j], "L") {
				l, _ := strconv.Atoi(strings.Replace(wires[i][j], "L", "", 1))
				rightLeft[i] -= l
			} else if strings.Contains(wires[i][j], "U") {
				u, _ := strconv.Atoi(strings.Replace(wires[i][j], "U", "", 1))
				upDown[i] += u
			} else if strings.Contains(wires[i][j], "D") {
				d, _ := strconv.Atoi(strings.Replace(wires[i][j], "D", "", 1))
				upDown[i] -= d
			}
			coordinates = append(coordinates, Point{x: rightLeft[i], y: upDown[i]})
		}
		wireCoordinates = append(wireCoordinates, Wire{coordinates: coordinates})
	}

	common := commonPoints(wireCoordinates[0], wireCoordinates[1])
	fmt.Println("Common Points:")
	fmt.Println(common)

}

func commonPoints(wire1 Wire, wire2 Wire) (commonPoints []Point) {
	for i := 0; i < len(wire1.coordinates); i++ {
		for j := i; j < len(wire2.coordinates); j++ {
			if wire1.coordinates[i] == wire2.coordinates[j] {
				commonPoints = append(commonPoints, wire1.coordinates[i])
			}
		}
	}
	return commonPoints
}
