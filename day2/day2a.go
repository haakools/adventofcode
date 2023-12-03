package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
)

func read_file_to_str(path string) string {
	body, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Error reading the file %s", err)
	}
	str := string(body)
	return str
}

func parseLine(line string) []string {
	// split on ":"
	// then split on ";" and ","
	colon_split := strings.Split(line, ":")
	//fmt.Printf("type %T\n", colon_split[1])

	each_try := strings.Split(colon_split[1], ";")
	return each_try
}

func isNumber(s string) bool {
	var isNum bool; // defaults to false
	if _, err := strconv.Atoi(s); err == nil {
		isNum = true
	}
	return isNum
}
// Do it with runes and maps instead of this ugly string stuff

// Find the digit

func validateGame(samples []string, bagContents map[string]int) bool {
	for _, sample := range(samples) {
		for _, cubeInfo := range(strings.Split(sample, ",")) {

			cubeInfoParsed := strings.Split(cubeInfo, " ")
			fmt.Println(cubeInfoParsed)
			for _, item := range(cubeInfoParsed) {
				fmt.Println(bagContents[item])
			}
			number, _ := strconv.Atoi(cubeInfoParsed[1]);
			color := cubeInfoParsed[2]
			if color == "" {
				break
			}
			//fmt.Printf("The parsed color is %v\n", color)

			countDifference := bagContents[color] - number
			fmt.Printf("Color %v has count %v. Difference is %d\n", color, number, countDifference)
			if countDifference < 0 {
				return false
			}
		}

	}
	return true
}

func main() {
	string_data := read_file_to_str("example.txt")
	lines := strings.Split(string_data, "\n")

	bagContents := make(map[string]int)
	bagContents["red"]   = 12
	bagContents["green"] = 13
	bagContents["blue"]  = 14

	validCounter := 0

	for i:= 0; i < len(lines); i++ {
		fmt.Printf("%v\n", lines[i])
		samples := parseLine(lines[i])

		if validateGame(samples, bagContents) {
			fmt.Printf("%d is valid\n", i)
			validCounter++
		}
		if i==0 {break}
	}


	fmt.Printf("There are %d valid games\n", validCounter)
	
}
