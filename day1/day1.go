package main

import (
	"fmt"
	"strings"
	"strconv"
	"log"
	"os"
)
func check(e error) {
	if e != nil {
		log.Fatalf("unable to read file: %v", e)
	}
}

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

func find_last_digit(line_str string, first_digit_idx int) string {
	var last_digit_str string = "0";

	for j:=len(line_str)-1; j>=first_digit_idx; j-- {
		if isNumber(line_str[j]) {
			last_digit_str = string(line_str[j])
			break;
		}
	}
	return last_digit_str
}

func main() {
	string_data := read_file_to_str("input.txt")
	lines := strings.Split(string_data, "\n")
	var result int = 0;

	for i:= 0; i < len(lines); i++ {
		fmt.Println(lines[i])
		first_digit, idx := find_first_digit(lines[i])
		last_digit := find_last_digit(lines[i], idx)
		number_as_str := first_digit + last_digit;

		number, err := strconv.Atoi(number_as_str);
		if err!= nil {
			fmt.Printf("Error reading the file %s", err)
		}
		
		fmt.Println(number)

		result = result + number;
	}
	fmt.Printf("The sum if %d\n", result);


}
