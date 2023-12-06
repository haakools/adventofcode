package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type NumLoc struct {
	start int
	end   int
}

type coord struct {
	x int
	y int
}

func isDot(r rune) bool {
	return r == '.'
}
func stringToNumber(s string) int {
	num, err := strconv.Atoi((s))
	if err != nil {
		fmt.Println("str conversion failed", err)
	}
	return num
}

func sum(slice []int) int {
	result := 0
	for _, val := range slice {
		result += val
	}
	return result
}

var length = 10
var width = 10

func readFileToRunes(path string) ([][]rune, [][]NumLoc) {

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	length := 10 // this should be read instead of constant
	grid := make([][]rune, length)
	line_matches := make([][]NumLoc, length)
	i := 0
	for scanner.Scan() {
		pattern := regexp.MustCompile("[0-9]+")
		matches := pattern.FindAllStringSubmatchIndex(scanner.Text(), -1) // matches is [][]string

		for _, match := range matches {
			line_matches[i] = append(line_matches[i], NumLoc{match[0], match[1]})
		}

		grid[i] = []rune(scanner.Text())
		i++
	}

	return grid, line_matches
}

func getOuterCoordinates(gridpoint NumLoc, lineNumber int) []coord {
	// Cannot access out-of-range values in the grid slice
	// So checking if they are valid here
	var output []coord
	// Append the points to the output
	for i := lineNumber - 1; i < lineNumber+1; i++ {
		if i >= 0 && i < length {
			for j := gridpoint.start - 1; j < gridpoint.end+1; j++ {
				if j >= 0 && j < width {
					output = append(output, coord{i, j})
				}

			}
		}

	}
	// i between [NumLoc.start-1, NumLoc.end+1]
	// [NumLoc.start-1, NumLoc.end+1]
	return output
}

func validatePoint(coords []coord, grid [][]rune) bool {
	for _, c := range coords {
		//grid_val := grid[c.start:c]
		grid_val := grid[c.x][c.y]
		//fmt.Println("Grid value", grid_val)
		fmt.Println(c.x, c.y)
		if !isDot(grid_val) || !unicode.IsDigit(grid_val) {
			return false
		}
	}
	return true
}

func main() {

	grid, gridpoints := readFileToRunes("example.txt")

	var gear_numbers []int

	for i := range grid {

		if len(gridpoints[i]) != 0 {
			// Create a slice of coords to check. The slice contains a list of
			for _, index := range gridpoints[i] {
				coords := getOuterCoordinates(index, i)
				fmt.Println(coords)
				if validatePoint(coords, grid) {
					gear_num := stringToNumber(string(grid[i][index.start:index.end]))
					gear_numbers = append(gear_numbers, gear_num)
				}
			}
		}

	}
	gearSum := sum(gear_numbers)
	fmt.Println("The sum is", gearSum)

}
