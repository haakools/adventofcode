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
	colon_split := strings.Split(line, ":")
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
	color := strings.TrimSpace(cubeInfo[strings.LastIndex(cubeInfo, " ")+1:])
	return number, color
}

func getLowestSet(samples []string) (int, int, int) {

	setCount := make(map[string]int)
	setCount["red"]	  	= 0
	setCount["blue"]	= 0	
	setCount["green"]	= 0

	for _, sample := range(samples) {
		colorCountStr := strings.Split(sample, ",")		
		for _, cubeInfo := range(colorCountStr) {
			number, color := splitSample(cubeInfo)

			if setCount[color] < number {
				setCount[color] = number
			}
		}
	}
	r := setCount["red"]
	b := setCount["blue"]
	g := setCount["green"]

	return r, b, g
}

func main() {
	string_data := read_file_to_str("input.txt")
	lines := strings.Split(string_data, "\n")

	setProductSum := 0

	for i:= 0; i < len(lines); i++ {
		fmt.Printf("%v\n", lines[i])
		samples := parseLine(lines[i])
		r,b,g := getLowestSet(samples) 
		setProductSum += r*b*g;
	}
	fmt.Printf("Set product sum is %d ", setProductSum)
}
