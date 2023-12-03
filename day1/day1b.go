package main

import (
	"fmt"
	"strings"
	"strconv"
	"sort"
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

func isNumber(s byte) bool {
	var isNum bool; // defaults to false
	if _, err := strconv.Atoi(string(s)); err == nil {
		isNum = true
	}
	return isNum
}

func find_first_digit(line_str string) (string, int) {
	var first_digit_str string = "0";
	var idx int = 0;

	for i:=0; i<len(line_str); i++ {
		if isNumber(line_str[i]) {
			// Converting because it is a byte
			first_digit_str = string(line_str[i])
			idx = i;
			break;
		}
	}
	return first_digit_str, idx
}

func find_last_digit(line_str string, first_digit_idx int) (string, int) {
	var last_digit_str string = "0";
	var idx int = 0;
	for j:=len(line_str)-1; j>=first_digit_idx; j-- {
		if isNumber(line_str[j]) {
			last_digit_str = string(line_str[j])
			idx = j;
			break;
		}
	}
	return last_digit_str, idx
}

func findStringOccurences(line_str string, string_map map[string]int) (map[int]string, []int) {
	// create a map of the numbers and a sorted list of the locations
	substrMap := make(map[int]string)
	substrKeys := make([]int, 0)
	for key, val := range(string_map) {
		if strings.Contains(line_str, key) {
			//fmt.Printf("Found the number %v in %v\n", line_str, key)
			idx := strings.Index(line_str, key)
			substrMap[idx] = strconv.Itoa(val)
			substrKeys = append(substrKeys, idx)
		}
	}
	sort.Ints(substrKeys)
	return substrMap, substrKeys
}

func getOutputNumbers(line string, m map[string]int) (string, string) {
	substrMap, substrKeysSorted := findStringOccurences(line, m)
	firstLiteralDigit, firstLiteralIndex := find_first_digit(line)
	lastLiteralDigit, lastLiteralIndex := find_last_digit(line, firstLiteralIndex)

	var firstDigit, lastDigit string = "0", "0";

	// If no digits was found, proceed with the solution part a
	if len(substrKeysSorted) == 0 {
		return firstLiteralDigit, lastLiteralDigit		
	}
	// Checking if the first or the last value is
	// found before/after the literal digit.
	substrIdxFirst := substrKeysSorted[0]
	substrIdxLast := substrKeysSorted[len(substrKeysSorted)-1]
	if firstLiteralIndex < substrIdxFirst   {
		firstDigit = firstLiteralDigit;
	} else {
		firstDigit = substrMap[substrIdxFirst]
	}
	if lastLiteralIndex > substrIdxLast {
		lastDigit = lastLiteralDigit;
	} else {
		lastDigit = substrMap[substrIdxLast]
	}
	return firstDigit, lastDigit

}

func main() {
	string_data := read_file_to_str("input.txt")
	lines := strings.Split(string_data, "\n")

	// creating a map 
	m := make(map[string]int)
	m["zero"] = 0
	m["one"]  = 1
	m["two"]  = 2
	m["three"]= 3
	m["four"] = 4
	m["five"] = 5
	m["six"]  = 6
	m["seven"]= 7
	m["eight"]= 8
	m["nine"] = 9

	var result int = 0;

	for i:= 0; i < len(lines); i++ {
		fmt.Println(lines[i])

		firstDigit, lastDigit := getOutputNumbers(lines[i], m)

		fmt.Printf("Number: %s%s\n", firstDigit, lastDigit)

		// Converting and summing the numbers	
		number, err := strconv.Atoi((firstDigit + lastDigit));

		if err!= nil {
			fmt.Printf("Error reading the file %s", err)
		}
		//if i == 5 {
		//	break
		//}
		result = result + number;
	}
	fmt.Printf("The sum if %d\n", result);
}
