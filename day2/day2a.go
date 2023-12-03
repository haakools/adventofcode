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

	samples := strings.Split(colon_split[1], ";")
	return samples
}

func isNumber(s string) bool {
	var isNum bool; // defaults to false
	if _, err := strconv.Atoi(s); err == nil {
		isNum = true
	}
	return isNum
}

func splitSample(cubeInfo string) (int, string) {
	cubeInfoParsed := strings.Split(cubeInfo, " ")
	number, _ := strconv.Atoi(cubeInfoParsed[1]);
	color_last_idx := strings.LastIndex(cubeInfo, " ")
	color := cubeInfo[color_last_idx+1:]
	return number, color
}


func validateGame(samples []string, bagContents map[string]int) bool {
	for _, sample := range(samples) {
		colorCountStr := strings.Split(sample, ",")		

		for _, cubeInfo := range(colorCountStr) {
			number, color := splitSample(cubeInfo)
			countDifference := bagContents[strings.TrimSpace(color)] - number
			if countDifference < 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	string_data := read_file_to_str("input.txt")
	lines := strings.Split(string_data, "\n")

	bagContents := make(map[string]int)
	bagContents["red"]   = 12
	bagContents["green"] = 13
	bagContents["blue"]  = 14

	IDsum := 0

	for i:= 0; i < len(lines); i++ {
		fmt.Printf("%v\n", lines[i])
		samples := parseLine(lines[i])

		if validateGame(samples, bagContents) {
			fmt.Println("GAME IS VALID.")
			IDsum += i+1
		}
	}


	fmt.Printf("The sum of valid IDs are %d \n", IDsum)
	
}
