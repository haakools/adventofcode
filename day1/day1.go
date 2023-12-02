package main


import (
	"fmt"
	"log"
	"os"
)
func check(e error) {
	if e != nil {
		log.Fatalf("unable to read file: %v", e)
	}
}

func read_file(path string) string {
	body, err := os.ReadFile(path)
	check(err)
	str := string(body)
	return str
}


func main() {
	string_data := read_file("example.txt")
	fmt.Println(string_data)
}
